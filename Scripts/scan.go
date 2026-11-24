package main

import (
	"fmt"
	"net"
	"sort"
	"sync"
)

func Scan(cfg Config, ports []int) []string {
	if len(ports) == 0 {
		return nil
	}

	workers := cfg.Workers
	if workers <= 0 {
		workers = 1
	}
	if workers > len(ports) {
		workers = len(ports)
	}

	jobs := make(chan int)
	results := make([]string, 0)
	var wg sync.WaitGroup
	var mu sync.Mutex

	wg.Add(workers)
	for range workers {
		go func() {
			defer wg.Done()
			for port := range jobs {
				address := fmt.Sprintf("%v:%v", cfg.Target, port)
				conn, err := net.DialTimeout("tcp", address, cfg.Timeout)
				if err != nil {
					continue
				}

				bannerText := banner(conn, cfg.Timeout)
				_ = conn.Close()

				var line string
				if bannerText != "" {
					line = fmt.Sprintf("[*] Port %v is Opened -> %s", port, bannerText)
					fmt.Printf("%v%s%v\n", Green, line, Reset)
				} else {
					service, ok := serviceMap[port]
					if !ok {
						service = "UNKNOWN"
					}
					line = fmt.Sprintf("[*] Port %v is Opened -> %v", port, service)
					fmt.Printf("%v%s%v\n", Green, line, Reset)
				}

				mu.Lock()
				results = append(results, line)
				mu.Unlock()
			}
		}()
	}

	for _, port := range ports {
		jobs <- port
	}
	close(jobs)
	wg.Wait()
	return results
}

func loop(cfg Config) {
	ports := knownPorts()
	Scan(cfg, ports)
}

func knownPorts() []int {
	ports := make([]int, 0, len(serviceMap))
	for port := range serviceMap {
		ports = append(ports, port)
	}
	sort.Ints(ports)
	return ports
}

func popularPorts() []int {
	const limit = 20
	if len(servicePorts) < limit {
		return append([]int(nil), servicePorts...)
	}
	return append([]int(nil), servicePorts[:limit]...)
}

func ScanAll(cfg Config) []string {
	ports := make([]int, 65535)
	for i := range ports {
		ports[i] = i + 1
	}
	return Scan(cfg, ports)
}
