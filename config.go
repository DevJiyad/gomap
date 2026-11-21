package main

import (
	"flag"
	"fmt"
	"time"
)

// =========================================================================
// HOW TO ADD A NEW FLAG LATER (3 EASY STEPS):
//   STEP 1: Add a new field to the Config struct below (e.g., Top bool)
//   STEP 2: Bind the flag inside ParseFlags() using &cfg.YourField
//   STEP 3: Use cfg.YourField inside main.go or Scan.go!
// =========================================================================

// Config is our "flag container" that holds all terminal inputs and settings.
type Config struct {
	Target  string        // Positional argument: the IP or domain to scan (flag.Arg(0))
	Port    int           // -p flag: single port to scan
	All     bool          // -all flag: true if user wants to scan all 65,535 ports
	Workers int           // -w flag: max concurrent goroutines scanning at once
	Timeout time.Duration // -t flag: how long to wait before a port is marked closed

	// [STEP 1] Add new flag variables right here:
}

// Usage prints custom red error text and instructions when a user types a bad command.
func Usage() {
	fmt.Println(Red + "Invalid command." + Reset)
	fmt.Println("Usage:")
	fmt.Println("  ./gomap -p <port> <target>   (Scan a single port)")
	fmt.Println("  ./gomap -all <target>        (Scan all 65,535 ports)")
}

// ParseFlags reads the user's terminal command and fills our Config container.
func ParseFlags() Config {
	var cfg Config
	flag.Usage = Usage

	// [STEP 2] Connect terminal flags to your Config struct using pointers (&):
	flag.IntVar(&cfg.Port, "p", 0, "Single port number to scan")
	flag.BoolVar(&cfg.All, "all", false, "Scan all 65,535 ports")
	flag.IntVar(&cfg.Workers, "w", 1000, "Number of concurrent workers")
	flag.DurationVar(&cfg.Timeout, "t", 500*time.Millisecond, "Timeout per port (e.g. 500ms, 1s)")

	// Read all flags from the terminal (MUST happen before reading flag.Arg)
	flag.Parse()

	// Grab the target IP/domain (the first word without a '-' in front of it)
	cfg.Target = flag.Arg(0)

	return cfg
}
