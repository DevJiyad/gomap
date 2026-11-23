## Warninng

if you ./gomap or gomap and you just found something similar to this
``

---

/ \_**\_// ** \| \/ | / \ | \_ \

| | _ | | | | |\/| | / _ \ | |_ |
| |_| || |**| | | | |/ \_** \| **/
\_\_**| \__\_\_/|_| |_/_/ \_\_|

[*] Target: scanme.nmap.org | Timeout: 1s

[!] Scan finished in 1.001s
``

Do not panic just rerun the command

## Download & Install (Windows)

1. Download **`gomap.exe`** from the [Releases](https://github.com) page.
2. Run it directly in PowerShell or CMD from your Downloads folder:
   ```bash
   .\gomap.exe -p 80 scanme.nmap.org
   ```

### Optional: Run `gomap` From Any Folder (Add to PATH)

To use `gomap` globally from any directory in your terminal:

1. Move `gomap.exe` to a permanent folder (for example, `C:\Tools`).
2. Press the **Windows Key**, search **Edit the system environment variables**, and click **Environment Variables**.
3. Under **User variables**, select **Path** → **Edit** → **New**, and paste `C:\Tools`.
4. Open a new terminal window and run:
   ```bash
   gomap -p 80 scanme.nmap.org
   ```

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

# Scan a comma-separated list or range of ports
./gomap.exe -p 22,80-90,443 scanme.nmap.org

# Scan all 65,535 ports
./gomap.exe -all scanme.nmap.org

# Custom timeout and worker count
./gomap.exe -all -w 1500 -t 300ms scanme.nmap.org
```
