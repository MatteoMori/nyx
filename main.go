package main

import (
	"os"
	"os/signal"
	"syscall"

	nyx "github.com/MatteoMori/nyx/cmd"
)

func main() {
	// Set up channel to listen for interrupt or terminate signals for graceful shutdown
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	// Run nyx.Execute in a goroutine so we can listen for signals concurrently
	done := make(chan struct{})
	go func() {
		nyx.Execute() // Execute the root CLI command (./cmd/nyx/root.go)
		close(done)
	}()

	select {
	case <-sigs:
		// Handle graceful shutdown here if needed
		// For example, we could call a shutdown function in nyx package
	case <-done:
		// nyx.Execute finished execution
	}
}
