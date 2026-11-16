package main

import (
	"fmt"
	"net"
	"sync"
	"time"
)

func main() {
	yellow := "\033[1;33m"
	green := "\033[1;32m"
	cyan := "\033[1;36m"
	reset := "\033[0m"
	target := "scanme.nmap.org"
	var port int
	banner := yellow + `
====================================================================
   ____  ____  ____  _____   ____   ____    _    _   _ _   _ _____ ____  

  |  _ \/ __ \|  _ \|_   _| / ___| / ___|  / \  | \ | | \ | | ____|  _ \ 
  | |_) | |  | | |_) | | |   \___ \| |     / _ \ |  \| |  \| |  _| | |_) |
  |  __/| |__| |  _ <  | |    ___) | |___ / ___ \| |\  | |\  | |___|  _ < 
  |_|    \____/|_| \_\ |_|   |____/ \____/_/   \_\_| \_|_| \_|_____|_| \_\
====================================================================
  Fast Concurrent TCP Port Scanner | Open-Source Edition
  Guide:
    - Scans TCP ports 1-65535 concurrently using Go routines
    - Identifies key pentesting services (FTP, SSH, HTTP, HTTPS)
    - Provides quick enumeration tips for discovered open ports
  Legal Notice: Only scan targets you own or have permission to test.
====================================================================` + reset

	fmt.Println(banner)

	// Resolve IP once before the loop so Go doesn't do 65,535 DNS lookups
	ips, err := net.LookupIP(target)
	if err != nil || len(ips) == 0 {
		fmt.Printf("[!] Could not resolve target: %v\n", target)
		return
	}
	targetIP := ips[0].String()

	fmt.Printf("%s[*] Target: %s (%s)%s\n", yellow, target, targetIP, reset)
	start := time.Now()
	fmt.Printf("%v[*] Scanning %v", yellow, reset)
	var wg sync.WaitGroup
	limit := make(chan struct{}, 3000) // Higher concurrency for much faster scanning
	for port = 1; port <= 65535; port++ {
		limit <- struct{}{}
		wg.Go(func() {
			defer func() { <-limit }()

			address := fmt.Sprintf("%s:%d", targetIP, port)
			// 350ms timeout cuts wait time by nearly 3x compared to 1 second
			con, err := net.DialTimeout("tcp", address, 350*time.Millisecond)
			if err != nil {
				return
			}
			defer con.Close()

			switch port {
			case 20, 21:
				fmt.Printf("%s[+] Port %d is open%s -> %sFTP%s\n", green, port, reset, cyan, reset)
			case 22, 2222:
				fmt.Printf("%s[+] Port %d is open%s -> %sSSH%s\n", green, port, reset, cyan, reset)
			case 80, 8000, 8008, 8080, 8888:
				fmt.Printf("%s[+] Port %d is open%s -> %sHTTP%s\n", green, port, reset, cyan, reset)
			case 443, 8443:
				fmt.Printf("%s[+] Port %d is open%s -> %sHTTPS%s\n", green, port, reset, cyan, reset)
			default:
				fmt.Printf("%s[+] Port %d is open%s\n", green, port, reset)
			}
		})
	}

	wg.Wait()

	fmt.Printf("%s[*] Scan finished in %.2fs%s\n", yellow, time.Since(start).Seconds(), reset)

	footer := yellow + `
╔═══════════════════════════╗
║     Made by dev/jiyad     ║
╚═══════════════════════════╝` + reset

	fmt.Println(footer)
}
