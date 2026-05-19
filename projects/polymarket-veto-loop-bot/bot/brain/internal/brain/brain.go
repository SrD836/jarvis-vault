package brain

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/jarvis/polymarket-veto-loop-bot/bot/brain/internal/llmclient"
	"github.com/jarvis/polymarket-veto-loop-bot/bot/brain/internal/rules"
	"github.com/jarvis/polymarket-veto-loop-bot/bot/brain/internal/types"
)

const (
	candidatesPath = "vault/inbox/trading/candidates.jsonl"
	approvedPath   = "vault/inbox/trading/approved.jsonl"
	blockedPath    = "vault/inbox/trading/blocked.jsonl"
)

// Run reads candidates from JSONL, vets them against V1-V6, writes approved/blocked.
func Run() error {
	if err := ensureDir(approvedPath); err != nil {
		return fmt.Errorf("ensure output dir: %w", err)
	}

	candidates, err := readCandidates(candidatesPath)
	if err != nil {
		return fmt.Errorf("read candidates: %w", err)
	}

	log.Printf("Brain: processing %d candidates", len(candidates))

	approved := make([]types.Approved, 0)
	blocked := make([]types.VetoResult, 0)

	for _, c := range candidates {
		// 1. Numeric rules (V1, V2, V4) — deterministic, no LLM
		if vr := rules.EvaluateNumeric(&c); vr != nil {
			blocked = append(blocked, *vr)
			log.Printf("VETOED %s by %s: %s", c.Slug, vr.VetoedBy, vr.Reason)
			continue
		}

		// 2. Semantic rules (V3, V5, V6) — LLM auditor
		llmResult := llmclient.EvaluateV3V5V6(&c)
		if llmResult == nil {
			// LLM unreachable or parse error — log and pass (conservative)
			log.Printf("WARN: LLM returned nil for %s, passing", c.Slug)
		} else if llmResult.Block {
			blocked = append(blocked, types.VetoResult{
				CandidateID: c.ID,
				Slug:        c.Slug,
				Blocked:     true,
				Reason:      fmt.Sprintf("%s: %s", llmResult.Rule, llmResult.Reason),
				VetoedBy:    llmResult.Rule,
			})
			log.Printf("VETOED %s by %s: %s", c.Slug, llmResult.Rule, llmResult.Reason)
			continue
		}

		// Passed all vetoes
		approved = append(approved, types.Approved{
			CandidateID:      c.ID,
			Slug:             c.Slug,
			Question:         c.Question,
			Category:         c.Category,
			CurrentPriceYes:  c.CurrentPriceYes,
			Volume24h:        c.Volume24h,
			EndDate:          c.EndDate,
			ScannedAt:        c.ScannedAt,
			ApprovedAt:       time.Now().UTC().Format(time.RFC3339),
			ApprovedPriceYes: c.CurrentPriceYes,
		})
		log.Printf("APPROVED %s", c.Slug)
	}

	// Write outputs
	if err := writeJSONL(approvedPath, approved); err != nil {
		return fmt.Errorf("write approved: %w", err)
	}
	if err := writeJSONL(blockedPath, blocked); err != nil {
		return fmt.Errorf("write blocked: %w", err)
	}

	log.Printf("Brain done: %d approved, %d blocked", len(approved), len(blocked))
	return nil
}

func readCandidates(path string) ([]types.Candidate, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("open %s: %w", path, err)
	}
	defer f.Close()

	var candidates []types.Candidate
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		var c types.Candidate
		if err := json.Unmarshal([]byte(line), &c); err != nil {
			log.Printf("WARN: skip unparseable candidate line: %v", err)
			continue
		}
		candidates = append(candidates, c)
	}

	return candidates, scanner.Err()
}

func writeJSONL(path string, data interface{}) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()

	enc := json.NewEncoder(f)
	switch v := data.(type) {
	case []types.Approved:
		for _, item := range v {
			if err := enc.Encode(item); err != nil {
				return err
			}
		}
	case []types.VetoResult:
		for _, item := range v {
			if err := enc.Encode(item); err != nil {
				return err
			}
		}
	default:
		return fmt.Errorf("unsupported type for JSONL write")
	}

	return nil
}

func ensureDir(filePath string) error {
	dir := filepath.Dir(filePath)
	return os.MkdirAll(dir, 0755)
}
