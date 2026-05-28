package llmclient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/jarvis/polymarket-veto-loop-bot/bot/brain/internal/types"
)

const dashboardURL = "http://localhost:3000/api/llm/call"
const defaultModel = "claude-sonnet-4-6"

// SystemPromptEdgeGate is the v7 prompt. The LLM still vetoes V3/V5/V6 but is
// also forced to emit an explicit probability estimate, edge type, edge
// description and thesis-invalidation condition — these power the E1/E2 gates
// and the calibration registry.
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
- estimated_prob: tu probabilidad real (no la implícita del market) de que YES resuelva true. Rango [0.01, 0.99]. Calíbrate honestamente: si no tienes edge, ponla cerca del precio implícito.
- edge_type: "info" si tienes información asimétrica concreta; "arb" si hay desalineación con otro market; "calibration" si la categoría tiene sesgo histórico documentado; "liquidity" si haces market making; "other" si no encaja; "none" si no hay edge real (en ese caso el brain bloqueará con E1).
- edge_description: una frase concreta. Ejemplo bueno: "fuente Reuters 2026-05-26 confirma Trump anunciará Powell antes del 30". Ejemplo malo: "el precio parece bajo".
- thesis_invalidation: qué tendría que pasar para que cierres la posición. Ejemplo: "Trump retira la nominación o el Senado vota en contra antes del 30".

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
		return &types.LLMBlockResult{Block: false, Reason: fmt.Sprintf("LLM decode error: %v", err), Rule: ""}
	}

	var result types.LLMBlockResult
	if err := json.Unmarshal([]byte(llmResp.Content), &result); err != nil {
		return &types.LLMBlockResult{Block: false, Reason: fmt.Sprintf("LLM output parse error: %s", llmResp.Content), Rule: ""}
	}

	return &result
}
