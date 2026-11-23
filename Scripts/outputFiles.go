package main

import (
	"fmt"
	"os"
)

func Output(filename string, results []string, cfg Config) {
	if filename == "" {
		return
	}

	file, err := os.Create(filename)
	if err != nil {
		fmt.Printf("Error creating output file: %v\n", err)
		return
	}
	defer file.Close()

	if _, err := fmt.Fprintf(file, "Target: %s\n", cfg.Target); err != nil {
		fmt.Printf("Error writing output file: %v\n", err)
		return
	}
	for _, line := range results {
		if _, err := fmt.Fprintf(file, "%s\n", line); err != nil {
			fmt.Printf("Error writing output file: %v\n", err)
			return
		}
	}
}
