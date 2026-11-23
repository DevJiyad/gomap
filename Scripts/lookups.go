package main

import (
	"fmt"
	"net"
	"strings"
)

func Lookup(cfg *Config) {
	parseIP := net.ParseIP(cfg.Target)
	if parseIP == nil {
		ips, err := net.LookupIP(cfg.Target)
		if err != nil {
			fmt.Printf("Host %v could not be resolved\n", cfg.Target)
			return
		}
		cfg.Target = ips[0].String()
		fmt.Printf("[*] Resolved target to: %s\n", cfg.Target)
	}

	if cfg.All || cfg.PortSpec != "" {
		return
	}

	loop(*cfg)
}

func LookupDiscovery(cfg *Config) {
	parseIP := net.ParseIP(cfg.Target)
	if parseIP == nil {
		ips, err := net.LookupIP(cfg.Target)
		if err != nil {
			fmt.Printf("Host %v could not be resolved", cfg.Target)
			return
		}
		cfg.Target = ips[0].String()
	}
	conn, err := net.DialTimeout("tcp", net.JoinHostPort(cfg.Target, "80"), cfg.Timeout)
	if err != nil {
		if !strings.Contains(err.Error(), "refused") {
			fmt.Printf("Host %v is not running (unreachable)\n", cfg.Target)
			return
		}
	} else {
		conn.Close()
	}
	fmt.Printf("Host %v is running \n", cfg.Target)
}
