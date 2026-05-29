package postmortem

import (
	"strings"
	"testing"
)

// Guardrail: the loss post-mortem runs from the exit_monitor cron (*/5), so it
// must NEVER call Claude Max (automated Max usage is the account-ban risk the
// model-routing guardrail eliminates). It uses DeepSeek via the dashboard route.
// This test fails loudly if anyone reverts the model to claudemax/claude.
func TestPostmortemModelIsNotClaude(t *testing.T) {
	if strings.Contains(strings.ToLower(postmortemModel), "claude") {
		t.Fatalf("postmortem must NOT use Claude from cron: got %q", postmortemModel)
	}
	if postmortemModel != "deepseek/deepseek-chat" {
		t.Fatalf("expected deepseek/deepseek-chat, got %q", postmortemModel)
	}
}
