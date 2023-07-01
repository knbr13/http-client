package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/abdullah-alaadine/http-client/internal/commands"
	"github.com/gookit/color"
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

func printColoredHeaders(headers map[string][]string) {
	// Print headers with colors
	for key, values := range headers {
		for _, value := range values {
			color.Bold.Printf("%s: ", key)
			color.Cyan.Printf("%s\n", value)
		}
	}
}
