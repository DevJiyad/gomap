package main

import (
	"fmt"
	"net"
	"strings"
	"time"
)

func banner(conn net.Conn, timeout time.Duration) string {
	_ = conn.SetDeadline(time.Now().Add(timeout))
	buf := make([]byte, 1024)

	read, err := conn.Read(buf)
	if err != nil && read == 0 {
		if netErr, ok := err.(net.Error); ok && netErr.Timeout() {
			host, _, splitErr := net.SplitHostPort(conn.RemoteAddr().String())
			if splitErr != nil {
				host = "localhost"
			}

			_ = conn.SetDeadline(time.Now().Add(timeout))
			req := fmt.Sprintf("HEAD / HTTP/1.1\r\nHost: %s\r\nUser-Agent: gomap\r\nConnection: close\r\n\r\n", host)
			if _, writeErr := conn.Write([]byte(req)); writeErr != nil {
				return ""
			}

			read, err = conn.Read(buf)
			if err != nil && read == 0 {
				return ""
			}
		} else {
			return ""
		}
	}

	clean := strings.TrimSpace(string(buf[:read]))
	if clean == "" {
		return ""
	}

	// If it's an HTTP reply, keep only the status line, Date, Server, and X-Powered-By
	if strings.HasPrefix(strings.ToUpper(clean), "HTTP/") {
		var wanted []string
		for _, line := range strings.Split(clean, "\n") {
			line = strings.TrimSpace(line)
			lower := strings.ToLower(line)

			if strings.HasPrefix(lower, "http/") ||
				strings.HasPrefix(lower, "date:") ||
				strings.HasPrefix(lower, "server:") ||
				strings.HasPrefix(lower, "x-powered-by:") {
				wanted = append(wanted, line)
			}
		}
		if len(wanted) > 0 {
			return strings.Join(wanted, " | ")
		}
	}

	// For non-HTTP services (SSH, FTP, SMTP), keep only the first clean line
	firstLine := strings.Split(clean, "\n")[0]
	return strings.TrimSpace(firstLine)
}
