package main

import (
	"fmt"
	"os"

	"github.com/davidgnuez/polymarket-veto-bot/scanner/internal/scanner"
)

func main() {
	if err := scanner.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "[scanner] error: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("[scanner] done")
}
