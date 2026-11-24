package main

import (
	"flag"
	"fmt"
	"os"
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
	if serviceMapErr != nil {
		fmt.Fprintf(os.Stderr, "Warning: %v; service names will be UNKNOWN\n", serviceMapErr)
	}

	targetOnly := cfg.Target != "" && flag.NFlag() == 0
	noFlags := cfg.Target == "" || cfg.PortSpec == "" && !cfg.All && !cfg.Lookup && !cfg.LookupDiscovery && !cfg.Banner && !targetOnly

	if noFlags {
		fmt.Println(Red + "Error: Missing argument" + Reset)
		flag.Usage()
		return
	}

	var ports []int

	if targetOnly {
		ports = popPorts()
	}
	if cfg.PortSpec != "" && !cfg.All {
		var err error
		ports, err = ParsePorts(cfg.PortSpec)
		if err != nil {
			fmt.Println(Red + err.Error() + Reset)
			return
		}
	}
	if cfg.Banner && cfg.PortSpec == "" && !cfg.All {
		ports = popPorts()
	}

	PrintBanner(cfg)
	start := time.Now()

	var results []string
	if cfg.Lookup {
		Lookup(&cfg)
	}
	if cfg.LookupDiscovery {
		LookupDiscovery(&cfg)
	}

	if cfg.All {
		results = ScanAll(cfg)
	} else if len(ports) > 0 {
		results = Scan(cfg, ports)
	}

	fmt.Printf(Yellow+"\n[!] Scan finished in %v\n"+Reset, time.Since(start).Round(time.Millisecond))
	Output(cfg.Output, results, cfg)
}
