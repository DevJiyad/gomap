package main

import (
	"net"
	"strings"
	"time"
)

func banner(conn net.Conn, timeout time.Duration) string {
	_ = conn.SetReadDeadline(time.Now().Add(timeout))
	buf := make([]byte, 1024)
	read, err := conn.Read(buf)
	if err != nil && read == 0 {
		if netErr, ok := err.(net.Error); ok && netErr.Timeout() {
			_, _ = conn.Write([]byte("HEAD / HTTP/1.0\r\n\r\n"))
			_ = conn.SetReadDeadline(time.Now().Add(timeout))
			read, err = conn.Read(buf)
			if err != nil && read == 0 {
				return ""
			}
		} else {
			return ""
		}
	}
	clean := strings.TrimSpace(string(buf[:read]))

	// If it's an HTTP reply, keep only the status line, Date, and Server
	if strings.HasPrefix(clean, "HTTP/") {
		var wanted []string
		for _, line := range strings.Split(clean, "\n") {
			line = strings.TrimSpace(line)
			if strings.HasPrefix(line, "HTTP/") ||
				strings.HasPrefix(line, "Date:") ||
				strings.HasPrefix(line, "Server:") ||
				strings.HasPrefix(line, "X-Powered-By") {

			}
		}
		return strings.Join(wanted, " | ")
	}

	return clean
}
