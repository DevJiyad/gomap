# gomap

> **Note:** Built and tested primarily for Windows.

A fast, lightweight, concurrent TCP port scanner written in Go with zero external dependencies.

> **Warning / Disclaimer:** This tool is for educational purposes and authorized security testing only. Only scan systems and networks you own or have explicit permission to test.

## Features

- **Concurrent Scanning:** Uses goroutines and a worker pool (`-w`) to scan thousands of ports in under a second.
- **Banner Grabber (`-b`):** Connects to open ports to pull service banners (SSH, FTP, SMTP, etc.) and probes quiet HTTP/web servers automatically.
- **Built-in Service Detection:** Identifies 5,000+ common services using an embedded Nmap service database.
- **DNS & Quick Top-10 Scan (`-ip` / `-ipd`):** Resolve target domains to IP addresses with or without scanning the top 10 ports.
- **File Export (`-o`):** Save clean scan results directly to a text file.
- **Zero Dependencies:** Built purely with the Go standard library.

## Install via Go

```bash
go install github.com/DevJiyad/gomap/Scripts@latest
```

## Build from source (Developers)

```bash
git clone https://github.com/DevJiyad/gomap.git
cd gomap/Scripts
go build -o gomap.exe .
```

### Run `gomap` From Any Folder (Add to PATH)

To use `gomap` globally from any directory in your terminal:

1. Move `gomap.exe` to a permanent folder (for example, `C:\Tools`).
2. Press the **Windows Key**, search **Edit the system environment variables**, and click **Environment Variables**.
3. Under **User variables**, select **Path** → **Edit** → **New**, and paste `C:\Tools`.
4. Open a new terminal window and run:
   ```powershell
   gomap -p 80 scanme.nmap.org
   ```

```bash
# If you want to see all the commands
./gomap
```

## Examples

```bash
# Scan top ports with smart banner grabbing
gomap -b scanme.nmap.org

# Scan a custom port range with 1000 workers
gomap -p 20-100 -w 1000 scanme.nmap.org

# Save results to a file
gomap -p 22,80 -b -o results.txt scanme.nmap.org
```
