package main

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/abdullah-alaadine/http-client/internal/commands"
)

func main() {
	if len(os.Args) < 2 {
		printHelpMessage()
		return
	}

	httpCommands := commands.GetHttpCommands()
	command := os.Args[1]

	if httpCommand, ok := httpCommands[command]; ok {
		args := os.Args[2:]
		httpResponse, err := httpCommand.Run(args)
		if err != nil {
			log.Fatal(err)
		}

		// Handle the HTTP response based on the command
		switch command {
		case "httphead":
			// For HEAD request, only print the response status
			fmt.Println(formatHeaders(httpResponse.Header))
		default:
			// For other requests, print the response body
			defer httpResponse.Body.Close()
			body, err := io.ReadAll(httpResponse.Body)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(string(body))
		}
	} else {
		log.Fatalf("Invalid command: %s", command)
	}
}
