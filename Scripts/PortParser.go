package main

import (
	"fmt"
	"strconv"
	"strings"
)

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
