package main

import (
	"io"
	"net"
	"os"
	"path/filepath"
	"testing"
	"time"
)

func TestBannerReturnsPartialDataBeforeTimeout(t *testing.T) {
	client, server := net.Pipe()
	defer client.Close()
	defer server.Close()

	go func() {
		_, _ = io.WriteString(server, "220 greeting")
	}()

	got := banner(client, 20*time.Millisecond)
	if got != "220 greeting" {
		t.Fatalf("banner() = %q, want %q", got, "220 greeting")
	}
}

func TestLoadServicesReadsWorkingDirectoryFile(t *testing.T) {
	dir := t.TempDir()
	data := []byte("# test services\nsmtp 25/tcp 0.1\nhttp 80/tcp 0.3\nhttps 443/tcp 0.5\n")
	if err := os.WriteFile(filepath.Join(dir, "nmap-services"), data, 0600); err != nil {
		t.Fatal(err)
	}

	originalDir, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	if err := os.Chdir(dir); err != nil {
		t.Fatal(err)
	}
	defer func() {
		if err := os.Chdir(originalDir); err != nil {
			t.Errorf("restore working directory: %v", err)
		}
	}()

	services, ports, err := loadServices()
	if err != nil {
		t.Fatalf("loadServices() error = %v", err)
	}
	if services[25] != "SMTP" {
		t.Fatalf("service for port 25 = %q, want SMTP", services[25])
	}
	if len(ports) != 3 || ports[0] != 443 || ports[1] != 80 || ports[2] != 25 {
		t.Fatalf("ports = %v, want [443 80 25] ordered by frequency", ports)
	}
}

func TestPopularPortsLimitsToTwenty(t *testing.T) {
	originalPorts := servicePorts
	servicePorts = make([]int, 25)
	defer func() {
		servicePorts = originalPorts
	}()

	if got := len(popularPorts()); got != 20 {
		t.Fatalf("popularPorts() returned %d ports, want 20", got)
	}
}
