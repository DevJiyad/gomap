# gomap

A fast, lightweight, concurrent TCP port scanner written in Go with zero external dependencies.

> **Disclaimer:** This tool is for educational purposes and authorized security testing only.

## Features
- **Concurrent Scanning:** Uses goroutines and a worker pool (`-w`) to scan thousands of ports in seconds.
- **Built-in Service Detection:** Automatically identifies common services (SSH, HTTP, HTTPS, SMB, RDP, etc.) using an embedded port database.
- **Customizable Timeouts:** Control per-port connection timeouts (`-t`) for fast local scans or slower remote targets.
- **Zero Dependencies:** Built purely with the Go standard library.

## Installation & Build
```bash
git clone https://github.com/DevJiyad/gomap
cd gomap
go build -o gomap.exe .
```

## Usage
```bash
# Scan a single port
./gomap.exe -p 80 scanme.nmap.org

# Scan all 65,535 ports
./gomap.exe -all scanme.nmap.org

# Custom timeout and worker count
./gomap.exe -all -w 1500 -t 300ms scanme.nmap.org
```
