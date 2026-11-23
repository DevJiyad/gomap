package main

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
)

var serviceMap, servicePorts, serviceMapErr = loadServices()

func loadServices() (map[int]string, []int, error) {
	paths := []string{"nmap-services", filepath.Join("Scripts", "nmap-services")}
	if executable, err := os.Executable(); err == nil {
		dir := filepath.Dir(executable)
		paths = append(paths,
			filepath.Join(dir, "nmap-services"),
			filepath.Join(dir, "Scripts", "nmap-services"),
		)
	}

	var data []byte
	var lastErr error
	tried := make([]string, 0, len(paths))
	seen := make(map[string]bool)
	for _, path := range paths {
		if seen[path] {
			continue
		}
		seen[path] = true
		tried = append(tried, path)
		var err error
		data, err = os.ReadFile(path)
		if err == nil {
			break
		}
		lastErr = err
	}
	if lastErr != nil {
		return nil, nil, fmt.Errorf("could not load nmap-services (tried %s): %w", strings.Join(tried, ", "), lastErr)
	}

	services := make(map[int]string)
	frequencies := make(map[int]float64)
	for _, line := range strings.Split(string(data), "\n") {
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		fields := strings.Fields(line) // e.g. ["http", "80/tcp", "0.484143"]
		if len(fields) >= 2 && strings.HasSuffix(fields[1], "/tcp") {
			portStr := strings.TrimSuffix(fields[1], "/tcp")
			if port, err := strconv.Atoi(portStr); err == nil {
				services[port] = strings.ToUpper(fields[0])
				if len(fields) >= 3 {
					frequencies[port], _ = strconv.ParseFloat(fields[2], 64)
				}
			}
		}
	}

	ports := make([]int, 0, len(services))
	for port := range services {
		ports = append(ports, port)
	}
	sort.Slice(ports, func(i, j int) bool {
		if frequencies[ports[i]] == frequencies[ports[j]] {
			return ports[i] < ports[j]
		}
		return frequencies[ports[i]] > frequencies[ports[j]]
	})
	return services, ports, nil
}
