package scripts

import (
	"flag"
	"time"
)

// Config holds all terminal inputs and settings.
type Config struct {
	Target  string
	Port    int
	All     bool
	Lookup  bool
	Timeout time.Duration
}

// ParseFlags reads the user's terminal command and fills our Config container.
func ParseFlags() Config {
	var cfg Config

	flag.IntVar(&cfg.Port, "p", 0, "Single port number to scan")
	flag.BoolVar(&cfg.All, "all", false, "Scan all 65,535 ports")
	flag.BoolVar(&cfg.Lookup, "lp", false, "Lookup Ip address for target")
	flag.DurationVar(&cfg.Timeout, "timeout", 500*time.Millisecond, "Timeout per port (e.g. 500ms, 1s)")

	flag.Parse()
	cfg.Target = flag.Arg(0)

	return cfg
}
