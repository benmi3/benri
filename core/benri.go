package core

import (
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"

	"gitlab.com/benmi/benri/cmd" // Replace with your commands package
)

type Server struct {
	// Server configuration and state
}

func (s *Server) Start() error {
	// Start the server
	return nil
}

func (s *Server) Stop() error {
	// Stop the server
	return nil
}

func main() {
	server := NewServer() // Create a new Server instance

	// Start the server
	if err := server.Start(); err != nil {
		fmt.Println("Error starting server:", err)
		os.Exit(1)
	}

	// Command loop
	for {
		cmd := readCommand() // Read a command from the user

		// Process the command
		if err := commands.HandleCommand(server, cmd); err != nil {
			fmt.Println("Error processing command:", err)
		}
	}

	// Graceful shutdown
	server.Stop()
}
