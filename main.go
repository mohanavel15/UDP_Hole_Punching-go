package main

import (
	"fmt"
	"os"
)

func main() {
	cli_args := os.Args[1:]
	if len(cli_args) == 0 {
		fmt.Println("No arguments provided")
		os.Exit(1)
	}

	switch cli_args[0] {
	case "server":
		fmt.Println("Starting server")
		server()
	case "client":
		fmt.Println("Starting client")
		client(cli_args[1])
	default:
		fmt.Println("Invalid argument")
		os.Exit(1)
	}
}
