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

func scanSinglePort(target string, port int, timeout time.Duration) string {
	address := fmt.Sprintf("%v:%v", target, port)
	con, err := net.DialTimeout("tcp", address, timeout)
	if err != nil {
		return ""
	}
	defer con.Close()

	banner := grabBanner(con, timeout)
	if banner != "" {
		fmt.Printf("%v[*] Port %v is Opened -> %s%v\n", Green, port, banner, Reset)
		return fmt.Sprintf("[*] Port %v is Opened -> %s", port, banner)
	} else {
		fmt.Printf("%v[*] Port %v is Opened%v\n", Green, port, Reset)
		return fmt.Sprintf("[*] Port %v is Opened", port)
	}
}

func ScanPorts(cfg Config, ports []int) []string {
	var wg sync.WaitGroup
	sem := make(chan struct{}, cfg.Workers)
	var mu sync.Mutex
	results := make([]string, 0)

	for _, port := range ports {
		port := port
		wg.Add(1)
		go func() {
			defer wg.Done()
			sem <- struct{}{}
			defer func() { <-sem }()
			if line := scanSinglePort(cfg.Target, port, cfg.Timeout); line != "" {
				mu.Lock()
				results = append(results, line)
				mu.Unlock()
			}
		}()
	}
	wg.Wait()
	return results
}

func ScanAll(cfg Config) []string {
	var wg sync.WaitGroup
	sem := make(chan struct{}, cfg.Workers)
	var mu sync.Mutex
	results := make([]string, 0)

	for port := 1; port <= 65535; port++ {
		port := port
		wg.Add(1)
		go func() {
			defer wg.Done()
			sem <- struct{}{}
			defer func() { <-sem }()
			if line := scanSinglePort(cfg.Target, port, cfg.Timeout); line != "" {
				mu.Lock()
				results = append(results, line)
				mu.Unlock()
			}
		}()
	}
	wg.Wait()
	return results
}
