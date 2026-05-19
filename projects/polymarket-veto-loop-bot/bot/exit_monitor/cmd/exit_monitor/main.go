package main

import (
	"log"

	"github.com/davidgn/polymarket-veto-loop-bot/bot/exit_monitor/internal/monitor"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println("exit_monitor: starting one-shot pass")
	monitor.Run()
	log.Println("exit_monitor: done")
}
