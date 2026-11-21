package main

import (
	"fmt"
	"time"
)

const (
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Reset  = "\033[0m"
)

func main() {
	cfg := ParseFlags()

	if cfg.Target == "" || (cfg.Port == 0 && !cfg.All) {
		fmt.Println(Red + "Error: Missing port or target!" + Reset)
		Usage()
		return
	}

	start := time.Now()

	if cfg.All {
		ScanAll(cfg)
	} else {
		Scan(cfg)
	}

	fmt.Printf(Yellow+"\n[!] Scan finished in %v\n"+Reset, time.Since(start).Round(time.Millisecond))
}
