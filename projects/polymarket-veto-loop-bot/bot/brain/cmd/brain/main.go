package main

import (
	"log"

	"github.com/jarvis/polymarket-veto-loop-bot/bot/brain/internal/brain"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println("Brain starting...")

	if err := brain.Run(); err != nil {
		log.Fatalf("Brain failed: %v", err)
	}

	log.Println("Brain finished.")
}
