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
