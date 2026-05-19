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

// SystemPromptV3V5V6 instructs the LLM to evaluate V3, V5, V6.
const SystemPromptV3V5V6 = `Eres un auditor de mercados de predicción en Polymarket. 
Evalúas el mercado contra 3 reglas de veto. Devuelve SIEMPRE un JSON con la estructura
{"block": bool, "reason": string, "rule": string}.

Reglas:

V3 - Triggers vagos: El trigger textual del mercado no tiene fecha concreta o no tiene 
un referéndum/evento asociado verificable. Ejemplos de triggers vagos: "pronto", 
"en las próximas semanas", "cuando termine el proceso". Ejemplos de triggers concretos: 
"elecciones USA 2024", "decisión de la Fed 20 marzo", "lanzamiento producto X fecha Y".

V5 - Patrón débil: El evento tiene <3 fuentes independientes o no tiene precedente 
análogo histórico. Ejemplo débil: rumor sin fuentes. Ejemplo fuerte: hecho respaldado por 
3+ medios con precedente similar en otro mercado.

V6 - Sin trigger claro: No hay un catalyst identificable en los próximos 7 días con 
fecha de resolución concreta. Si la pregunta del mercado es vaga o no tiene un evento 
discreto asociado, esto es veto.

Responde solo el JSON. No añadas explicación fuera del JSON.`

// EvaluateV3V5V6 sends the candidate to the LLM for semantic veto evaluation.
// Returns the LLM result. If the dashboard is unreachable, returns a pass
// (conservative: don't block on infra failure).
func EvaluateV3V5V6(c *types.Candidate) *types.LLMBlockResult {
	payload := types.LLMRequest{
		Model:   defaultModel,
		System:  SystemPromptV3V5V6,
		Message: fmt.Sprintf("Market: %s\nCategory: %s\nQuestion: %s\nEnd date: %s\nVolume: %.0f\nPrice: %.4f",
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
