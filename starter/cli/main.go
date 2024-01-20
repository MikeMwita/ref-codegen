package main

import (
	"fmt"
	"os"
)

func main() {
	commands := NewCommands()

	// Check if there are any arguments
	if len(os.Args) < 2 {
		fmt.Println("No command provided. Use 'help' for available commands.")
		os.Exit(1)
	}

	// Get the first argument as the command
	command := os.Args[1]

	args := os.Args[2:]

	switch command {
	case "generate":
		commands.GenerateCommand(args)
	case "validate":
		commands.ValidateCommand(args)
	case "encrypt":
		commands.EncryptCommand(args)
	case "decrypt":

		commands.DecryptCommand(args)

	default:
		UnknownCommand()
	}
}
