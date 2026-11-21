package main

import (
	"fmt"
	"time"
)

const (
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Cyan   = "\033[36m"
	Reset  = "\033[0m"
)

func PrintBanner(cfg Config) {
	banner := `
   ______ ____  __  __    _    ____  
  / ____// __ \|  \/  |  / \  |  _ \ 

 | |  _ | |  | | |\/| | / _ \ | |_) |
 | |_| || |__| | |  | |/ ___ \|  __/ 
  \____| \____/|_|  |_/_/   \_\_|    
`
	fmt.Println(Cyan + banner + Reset)
	fmt.Printf(Yellow+"  [*] Target: %s | Timeout: %v | Workers: %d\n\n"+Reset, cfg.Target, cfg.Timeout, cfg.Workers)
}

func main() {
	cfg := ParseFlags()

	if cfg.Target == "" || (cfg.Port == 0 && !cfg.All) {
		fmt.Println(Red + "Error: Missing port or target!" + Reset)
		Usage()
		return
	}

	PrintBanner(cfg)
	start := time.Now()

	if cfg.All {
		ScanAll(cfg)
	} else {
		Scan(cfg)
	}

	fmt.Printf(Yellow+"\n[!] Scan finished in %v\n"+Reset, time.Since(start).Round(time.Millisecond))
}
