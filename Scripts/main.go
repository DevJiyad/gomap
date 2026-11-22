package scripts

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
	fmt.Printf(Yellow+"  [*] Target: %s | Timeout: %v\n\n"+Reset, cfg.Target, cfg.Timeout)
}

func main() {
	cfg := ParseFlags()

	noFlags := /*Required*/ cfg.Target == "" || /*Not Required*/ cfg.Port == 0 && !cfg.All && !cfg.Lookup

	if noFlags {
		fmt.Println(Red + "Error: Missing argument" + Reset)
		flag.Usage()
		return
	}

	PrintBanner(cfg)
	start := time.Now()

	if cfg.Lookup {
		LookupIP(cfg)
	}

	if cfg.All {
		ScanAll(cfg)
	} else if cfg.Port != 0 {
		Scan(cfg)
	}

	fmt.Printf(Yellow+"\n[!] Scan finished in %v\n"+Reset, time.Since(start).Round(time.Millisecond))
}
