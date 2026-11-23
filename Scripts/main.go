package main

import (
	"flag"
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
	fmt.Printf(Yellow+"  [*] Target: %s | Timeout: %v | Workers: %v\n\n"+Reset, cfg.Target, cfg.Timeout, cfg.Workers)
}

func main() {
	cfg := ParseFlags()

	noFlags := cfg.Target == "" || cfg.PortSpec == "" && !cfg.All && !cfg.Lookup && !cfg.LookupH

	if noFlags {
		fmt.Println(Red + "Error: Missing argument" + Reset)
		flag.Usage()
		return
	}

	var ports []int
	if cfg.PortSpec != "" && !cfg.All {
		var err error
		ports, err = ParsePorts(cfg.PortSpec)
		if err != nil {
			fmt.Println(Red + err.Error() + Reset)
			return
		}
	}

	PrintBanner(cfg)
	start := time.Now()

	if cfg.Lookup {
		LookupIP(cfg)
	}
	if cfg.LookupH {
		LookupHost(cfg)
	}

	var results []string
	if cfg.All {
		results = ScanAll(cfg)
	} else if cfg.PortSpec != "" {
		results = ScanPorts(cfg, ports)
	}

	fmt.Printf(Yellow+"\n[!] Scan finished in %v\n"+Reset, time.Since(start).Round(time.Millisecond))
	Output(cfg.Output, results, cfg)
}
