package main

import (
	"fmt"
	"net"
)

func LookupHost(cfg Config) {
	hs, err := net.LookupHost(cfg.Target)
	if err != nil {
		fmt.Printf("%s[!] Lookup failed for %s: %v%s\n", Red, cfg.Target, err, Reset)
		return
	}
	if len(hs) == 0 {
		fmt.Printf("%s[!] Lookup returned no Host for %s%s\n", Red, cfg.Target, Reset)
		return
	}

	fmt.Printf("%s[*] Host lookup for %s%s\n", Cyan, cfg.Target, Reset)
	seen := make(map[string]struct{}, len(hs))
	for _, h := range hs {
		if _, exists := seen[h]; exists {
			continue
		}
		seen[h] = struct{}{}
		fmt.Printf("%s%v    -> %s%s\n", Green, cfg.Target, h, Reset)
	}
	fmt.Printf("%vHost is running%v", Green, Reset)

}

func LookupIP(cfg Config) {
	ips, err := net.LookupIP(cfg.Target)
	if err != nil {
		fmt.Printf("%s[!] Lookup failed for %s: %v%s\n", Red, cfg.Target, err, Reset)
		return
	}
	if len(ips) == 0 {
		fmt.Printf("%s[!] Lookup returned no IP addresses for %s%s\n", Red, cfg.Target, Reset)
		return
	}

	fmt.Printf("%s[*] IP Lookup for %s%s\n", Cyan, cfg.Target, Reset)
	seen := make(map[string]struct{}, len(ips))
	for _, ip := range ips {
		address := ip.String()
		if _, exists := seen[address]; exists {
			continue
		}
		seen[address] = struct{}{}
		fmt.Printf("%s%v    -> %s%s\n", Green, cfg.Target, address, Reset)
	}
	fmt.Printf("%vIp is running%v", Green, Reset)
}
