# gomap

A fast, lightweight, concurrent TCP port scanner written in Go with zero external dependencies.

> **Warning / Disclaimer:** This tool is for educational purposes and authorized security testing only. Only scan systems and networks you own or have explicit permission to test.

## Features

- **Concurrent Scanning:** Uses goroutines and a worker pool (`-w`) to scan thousands of ports in seconds.
- **Built-in Service Detection:** Automatically identifies common services (SSH, HTTP, HTTPS, SMB, RDP, etc.) using an embedded port database.
- **Customizable Timeouts:** Control per-port connection timeouts (`-timeout`) for fast local scans or slower remote targets.
- **Zero Dependencies:** Built purely with the Go standard library.

### Run `gomap` From Any Folder (Add to PATH)

To use `gomap` globally from any directory in your terminal:

1. Move `gomap.exe` to a permanent folder (for example, `C:\Tools`).
2. Press the **Windows Key**, search **Edit the system environment variables**, and click **Environment Variables**.
3. Under **User variables**, select **Path** → **Edit** → **New**, and paste `C:\Tools`.
4. Open a new terminal window and run:
   ```powershell
   gomap -p 80 scanme.nmap.org
   ```

## Install via Go (Developers)

```bash
go install github.com/DevJiyad/gomap/Scripts
```

Or build from source:

```bash
git clone https://github.com/gomap/Scripts
cd gomap
go build -o gomap.exe .
```

## Usage

```bash
# Scan a single port
gomap -p 80 scanme.nmap.org

# Scan a list or range of ports
gomap -p 22,80,443 scanme.nmap.org
gomap -p 20-211 scanme.nmap.org

# Scan all 65,535 ports
gomap -all scanme.nmap.org

# Custom timeout and worker count
gomap -all -w 1500 -timeout 300ms scanme.nmap.org
```
