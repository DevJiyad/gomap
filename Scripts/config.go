package main

import (
	"flag"
	"fmt"
	"strconv"
	"strings"
	"time"
)

// Config holds all terminal inputs and settings.
type Config struct {
	Target   string
	PortSpec string
	All      bool
	Lookup   bool
	LookupH  bool
	Output   string
	Timeout  time.Duration
	Workers  int
}

// ParseFlags reads the user's terminal command and fills our Config container.
func ParseFlags() Config {
	var cfg Config

	flag.StringVar(&cfg.PortSpec, "p", "", "Ports to scan (e.g., 80, 22,80,443, or 20-211)")
	flag.BoolVar(&cfg.All, "all", false, "Scan all 65,535 ports")
	flag.BoolVar(&cfg.Lookup, "lp", false, "Lookup Ip address for target")
	flag.DurationVar(&cfg.Timeout, "timeout", 500*time.Millisecond, "Timeout per port (e.g. 500ms, 1s)")
	flag.BoolVar(&cfg.LookupH, "lph", false, "Lookup Host address for target")
	flag.StringVar(&cfg.Output, "o", "", "Save scan results to a file")
	flag.IntVar(&cfg.Workers, "w", 500, "Control the amount of workers")

	flag.Parse()
	cfg.Target = flag.Arg(0)

	if cfg.Workers <= 0 {
		fmt.Println("Can not make amount of workers negative zero or zero")
		fmt.Println("Changing amount of workers to 500")
		cfg.Workers = 500
	}

	return cfg
}

func ParsePorts(spec string) ([]int, error) {
	ports := make([]int, 0)
	seen := make(map[int]struct{})

	for _, item := range strings.Split(spec, ",") {
		item = strings.TrimSpace(item)
		if item == "" {
			return nil, fmt.Errorf("invalid port specification: empty port")
		}

		if strings.Contains(item, "-") {
			rangeParts := strings.Split(item, "-")
			if len(rangeParts) != 2 {
				return nil, fmt.Errorf("invalid port range %q: expected start-end", item)
			}

			start, err := strconv.Atoi(strings.TrimSpace(rangeParts[0]))
			if err != nil {
				return nil, fmt.Errorf("invalid start port in range %q", item)
			}
			end, err := strconv.Atoi(strings.TrimSpace(rangeParts[1]))
			if err != nil {
				return nil, fmt.Errorf("invalid end port in range %q", item)
			}
			if start < 1 || end > 65535 || start > end {
				return nil, fmt.Errorf("invalid port range %q: ports must satisfy 1 <= start <= end <= 65535", item)
			}

			for port := start; port <= end; port++ {
				if _, exists := seen[port]; !exists {
					ports = append(ports, port)
					seen[port] = struct{}{}
				}
			}
			continue
		}

		port, err := strconv.Atoi(item)
		if err != nil || port < 1 || port > 65535 {
			return nil, fmt.Errorf("invalid port %q: port must be between 1 and 65535", item)
		}
		if _, exists := seen[port]; !exists {
			ports = append(ports, port)
			seen[port] = struct{}{}
		}
	}

	if len(ports) == 0 {
		return nil, fmt.Errorf("invalid port specification: no ports provided")
	}
	return ports, nil
}
