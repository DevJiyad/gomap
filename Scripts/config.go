package main

import (
	"flag"
	"fmt"
	"time"
)

// Config holds all terminal inputs and settings.
type Config struct {
	Target          string
	PortSpec        string
	All             bool
	Lookup          bool
	Output          string
	Timeout         time.Duration
	Workers         int
	LookupDiscovery bool
	Banner          bool
}

// ParseFlags reads the user's terminal command and fills our Config container.
func ParseFlags() Config {
	var cfg Config

	flag.StringVar(&cfg.PortSpec, "p", "", "Ports to scan (e.g., 80, 22,80,443, or 20-211)")
	flag.BoolVar(&cfg.All, "all", false, "Scan all 65,535 ports")
	flag.BoolVar(&cfg.Lookup, "ip", false, "Lookup Ip address for target AND scans top 10 ports")
	flag.BoolVar(&cfg.LookupDiscovery, "ipd", false, "Lookup IP address for targets only without scanning top 10 ports")
	flag.DurationVar(&cfg.Timeout, "timeout", 500*time.Millisecond, "Timeout per port (e.g. 500ms, 1s)")
	flag.StringVar(&cfg.Output, "o", "", "Save scan results to a file")
	flag.IntVar(&cfg.Workers, "w", 200, "Control the amount of workers")
	flag.BoolVar(&cfg.Banner, "b", false, "Scan the 20 most common TCP ports and read server banners")
	flag.Parse()
	cfg.Target = flag.Arg(0)

	if cfg.Workers <= 0 {
		fmt.Println("Can not make amount of workers negative zero or zero")
		fmt.Println("Changing amount of workers to 1")
		cfg.Workers = 1
	}

	return cfg
}
