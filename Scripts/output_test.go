package main

import (
	"os"
	"path/filepath"
	"testing"
)

func TestOutputWritesCollectedResults(t *testing.T) {
	filename := filepath.Join(t.TempDir(), "results.txt")
	cfg := Config{Target: "127.0.0.1"}
	results := []string{
		"[*] Port 80 is Opened",
		"[*] Port 443 is Opened -> test banner",
	}

	Output(filename, results, cfg)

	got, err := os.ReadFile(filename)
	if err != nil {
		t.Fatalf("reading output file: %v", err)
	}
	want := "Target: 127.0.0.1\n" +
		"[*] Port 80 is Opened\n" +
		"[*] Port 443 is Opened -> test banner\n"
	if string(got) != want {
		t.Errorf("output file content = %q, want %q", got, want)
	}
}
