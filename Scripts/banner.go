package main

import (
	"bufio"
	"net"
	"strings"
	"time"
)

func banner(conn net.Conn, timeout time.Duration) string {
	_ = conn.SetReadDeadline(time.Now().Add(timeout))
	line, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil && len(line) == 0 {
		return ""
	}
	clean := strings.TrimSpace(line)
	return clean
}
