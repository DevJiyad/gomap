package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
	"sync"
	"time"
)

func grabBanner(con net.Conn, timeout time.Duration) string {
	con.SetReadDeadline(time.Now().Add(timeout))
	line, err := bufio.NewReader(con).ReadString('\n')
	if err != nil {
		return ""
	}
	return strings.TrimSpace(line)
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
		fmt.Printf("%s    -> %s%s\n", Green, address, Reset)
	}
}

func Scan(cfg Config) {
	address := fmt.Sprintf("%v:%v", cfg.Target, cfg.Port)
	con, err := net.DialTimeout("tcp", address, cfg.Timeout)
	if err != nil {
		return
	}
	defer con.Close()

	banner := grabBanner(con, cfg.Timeout)
	if banner != "" {
		fmt.Printf("%v[*] Port %v is Opened -> %s%v\n", Green, cfg.Port, banner, Reset)
	} else {
		fmt.Printf("%v[*] Port %v is Opened%v\n", Green, cfg.Port, Reset)
	}
}

func ScanAll(cfg Config) {
	var wg sync.WaitGroup

	for port := 0; port <= 65535; port++ {
		port := port
		wg.Go(func() {
			address := fmt.Sprintf("%v:%v", cfg.Target, port)
			con, err := net.DialTimeout("tcp", address, cfg.Timeout)
			if err != nil {
				return
			}
			defer con.Close()

			banner := grabBanner(con, cfg.Timeout)
			if banner != "" {
				fmt.Printf("%v[*] Port %v is Opened -> %s%v\n", Green, port, banner, Reset)
			} else {
				fmt.Printf("%v[*] Port %v is Opened%v\n", Green, port, Reset)
			}
		})
	}
	wg.Wait()
}
