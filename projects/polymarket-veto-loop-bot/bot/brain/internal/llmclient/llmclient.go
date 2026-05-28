package llmclient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/jarvis/polymarket-veto-loop-bot/bot/brain/internal/types"
)

const dashboardURL = "http://localhost:3000/api/llm/call"
const defaultModel = "claude-sonnet-4-6"

// SystemPromptEdgeGate is the v7 prompt. The LLM still vetoes V3/V5/V6 but is
// also forced to emit an explicit probability estimate, edge type, edge
// description and thesis-invalidation condition — these power the E1/E2 gates
// and the calibration registry.
//
// v7 P3 (2026-05-28): relaxed edge declaration. The previous prompt was so
// strict that the LLM defaulted to edge_type=none for 100% of candidates,
// blocking the bot entirely. Now public news from Reuters/AP/Bloomberg/NYT
// within the last 30 days counts as a valid "info" edge, and an "edge_type"
// concreto se exige en >=80% de los casos no-bloqueados.
const SystemPromptEdgeGate = `Eres un trader cuantitativo de mercados de predicción en Polymarket.
Operas con probabilidades implícitas de eventos binarios (resuelven 0 o 1), NO con activos continuos.

Devuelves SIEMPRE un único JSON con esta estructura exacta:
{
  "block": bool,
  "reason": string,
  "rule": string,
  "estimated_prob": float,
  "edge_type": "info" | "arb" | "calibration" | "liquidity" | "other" | "none",
  "edge_description": string,
  "thesis_invalidation": string
}

Vetos textuales (block=true):
- V3 Trigger vago: sin fecha concreta o sin evento verificable.
- V5 Patrón débil: <3 fuentes independientes o sin precedente análogo.
- V6 Sin catalyst: no hay evento discreto identificable en los próximos 7 días.

Si bloqueas, devuelve también tu mejor estimación de probabilidad y edge_type="none".

Si NO bloqueas:
- estimated_prob: tu probabilidad real (no la implícita del market) de que YES resuelva true. Rango [0.01, 0.99]. Calíbrate honestamente; si no tienes edge, ponla cerca del precio implícito.
- edge_type: declara un edge concreto SIEMPRE QUE PUEDAS. Sólo usa "none" cuando el market sea genuinamente eficiente y tú no veas ninguna ventaja informacional. Si hay cualquier información asimétrica, sesgo de calibración histórico documentado, o liquidez explotable, declara el tipo apropiado.
  · "info": tienes información concreta (PÚBLICA O PRIVADA) que el precio del market aún no ha incorporado. Fuentes públicas de Reuters / AP / Bloomberg / NYT / WSJ / FT del último mes cuentan como "info" válido si su contenido todavía no se refleja en el precio. NO necesitas información privada para declarar "info".
  · "arb": el mismo evento está mispriced en otro venue (Kalshi, Manifold, Metaculus, casas de apuestas tradicionales).
  · "calibration": la categoría/horizonte tiene sesgo histórico documentado (ej. mercados de geopolítica corto plazo systematically overprice rumores). Cita el dato concreto.
  · "liquidity": haces market making con spread positivo esperado.
  · "other": razón cuantitativa válida que no encaja en las 4 anteriores. Explica.
  · "none": el market es eficiente y NO tienes ventaja. El brain bloqueará con E1.
- edge_description: una frase concreta con la fuente o el dato. Ejemplo bueno: "Reuters 2026-05-26 confirma anuncio Trump-Powell antes del 30; precio aún en 0.34". Ejemplo malo: "el precio parece bajo".
- thesis_invalidation: qué tendría que pasar para cerrar la posición. Ejemplo: "Trump retira la nominación o el Senado vota en contra antes del 30".

Regla de calibración: si ≥80% de tus respuestas en una sesión son edge_type="none", probablemente estás siendo demasiado conservador. Re-evalúa.

NO añadas texto fuera del JSON.`

// SystemPromptV3V5V6 is retained as alias for any caller still referencing it.
// Deprecated: use SystemPromptEdgeGate.
const SystemPromptV3V5V6 = SystemPromptEdgeGate

// EvaluateV3V5V6 sends the candidate to the LLM for semantic veto evaluation.
// Returns the LLM result. If the dashboard is unreachable, returns a pass
// (conservative: don't block on infra failure).
func EvaluateV3V5V6(c *types.Candidate) *types.LLMBlockResult {
	payload := types.LLMRequest{
		Model:  defaultModel,
		System: SystemPromptEdgeGate,
		Message: fmt.Sprintf("Market: %s\nCategory: %s\nQuestion: %s\nEnd date: %s\nVolume: %.0f\nImplied price (YES): %.4f\n\nDevuelve el JSON con tu estimated_prob, edge_type, edge_description y thesis_invalidation.",
			c.Slug, c.Category, c.Question, c.EndDate, c.Volume24h, c.CurrentPriceYes),
	}

	body, err := json.Marshal(payload)
	if err != nil {
		return &types.LLMBlockResult{Block: false, Reason: fmt.Sprintf("LLM marshal error: %v", err), Rule: ""}
	}

	client := &http.Client{Timeout: 30 * time.Second}
	resp, err := client.Post(dashboardURL, "application/json", bytes.NewReader(body))
	if err != nil {
		return &types.LLMBlockResult{Block: false, Reason: fmt.Sprintf("LLM unreachable: %v", err), Rule: ""}
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return &types.LLMBlockResult{Block: false, Reason: fmt.Sprintf("LLM status %d", resp.StatusCode), Rule: ""}
	}

	var llmResp types.LLMResponse
	if err := json.NewDecoder(resp.Body).Decode(&llmResp); err != nil {
		return &types.LLMBlockResult{Block: false, Reason: fmt.Sprintf("LLM decode error: %v", err), Rule: "PARSE"}
	}

	// v7 P3: tolerate stray prose around the JSON object by slicing on first
	// '{' and last '}'. Some Claude responses wrap the JSON in a fenced block
	// or add a one-liner before it.
	raw := llmResp.Content
	jstart := strings.Index(raw, "{")
	jend := strings.LastIndex(raw, "}")
	if jstart < 0 || jend <= jstart {
		log.Printf("[LLM] %s | parse_fail raw=%q", c.Slug, truncateForLog(raw, 300))
		return &types.LLMBlockResult{Block: false, Reason: "LLM no JSON in output", Rule: "PARSE"}
	}
	var result types.LLMBlockResult
	if err := json.Unmarshal([]byte(raw[jstart:jend+1]), &result); err != nil {
		log.Printf("[LLM] %s | parse_fail err=%v raw=%q", c.Slug, err, truncateForLog(raw, 300))
		return &types.LLMBlockResult{Block: false, Reason: fmt.Sprintf("LLM output parse error: %v", err), Rule: "PARSE"}
	}

	log.Printf("[LLM] %s | block=%t rule=%s est_prob=%.3f edge=%s desc=%q",
		c.Slug, result.Block, result.Rule, result.EstimatedProb, result.EdgeType,
		truncateForLog(result.EdgeDescription, 80))
	return &result
}

func truncateForLog(s string, n int) string {
	s = strings.ReplaceAll(s, "\n", " ")
	if len(s) <= n {
		return s
	}
	return s[:n-3] + "..."
}
