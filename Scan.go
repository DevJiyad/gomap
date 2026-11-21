package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"net"
	"strconv"
	"strings"
	"sync"
)

// go:embed ports.txt
var embeddedPorts string

var portMap = make(map[int]string)

// Load embedded ports
func init() {
	scanner := bufio.NewScanner(strings.NewReader(embeddedPorts))
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		parts := strings.Fields(line)
		if len(parts) >= 2 {
			port, err := strconv.Atoi(parts[0])
			if err == nil {
				portMap[port] = parts[1]
			}
		}
	}
}

func getService(port int) string {
	if service, exists := portMap[port]; exists {
		return service
	}
	return "UNKNOWN"
}

// 1. Scans ONE single port (and tells you if it's closed too!)
func Scan(cfg Config) {
	address := fmt.Sprintf("%s:%d", cfg.Target, cfg.Port)
	col, err := net.DialTimeout("tcp", address, cfg.Timeout)
	if err != nil {
		fmt.Printf(Red+"[-] Port: %d is CLOSED or FILTERED\n"+Reset, cfg.Port)
		return
	}
	defer col.Close()

	service := getService(cfg.Port)
	fmt.Printf(Green+"[*] Port: %d is OPENED -> "+Yellow+"%s\n"+Reset, cfg.Port, service)
}

// 2. Scans ALL 65,535 ports concurrently
func ScanAll(cfg Config) {
	ips, err := net.LookupHost(cfg.Target)
	if err != nil || len(ips) == 0 {
		fmt.Println(Red + "Error: Could not resolve target host!" + Reset)
		return
	}
	targetIP := ips

	var wg sync.WaitGroup
	sem := make(chan struct{}, cfg.Workers)

	for port := 1; port <= 65535; port++ {
		port := port
		wg.Add(1)
		sem <- struct{}{}

		go func() {
			defer wg.Done()
			defer func() { <-sem }()

			address := fmt.Sprintf("%s:%d", targetIP, port)
			col, err := net.DialTimeout("tcp", address, cfg.Timeout)
			if err != nil {
				return
			}
			defer col.Close()

			service := getService(port)
			fmt.Printf(Green+"[*] Port: %d is OPENED -> "+Yellow+"%s\n"+Reset, port, service)
		}()
	}

	wg.Wait()
}
