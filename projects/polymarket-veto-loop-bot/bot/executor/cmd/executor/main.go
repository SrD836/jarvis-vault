package main

import (
	"log"

	"github.com/davidgn/polymarket-veto-loop-bot/bot/executor/internal/executor"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println("executor: starting one-shot pass")
	executor.Run()
	log.Println("executor: done")
}
