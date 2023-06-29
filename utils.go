package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/abdullah-alaadine/http-client/internal/commands"
)

func formatHeaders(headers http.Header) string {
	var formattedHeaders strings.Builder

	for key, values := range headers {
		formattedHeaders.WriteString(fmt.Sprintf("%s: %s\n", key, strings.Join(values, ", ")))
	}

	return formattedHeaders.String()
}

func printHelpMessage() {
	httpCommands := commands.GetHttpCommands()

	helpMessage := `
Usage: http-client [command] [arguments]

Commands:`

	for _, httpCommand := range httpCommands {
		helpMessage += fmt.Sprintf("\n  %s\t%s", httpCommand.Name, httpCommand.Description)
	}

	helpMessage += `
	
Example:
  http-client httpget http://example.com
  http-client httppost http://example.com "body={key1: value, key2: value}"
`

	fmt.Println(helpMessage)
}
