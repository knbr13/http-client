package main

import (
	"io"
	"log"

	"github.com/abdullah-alaadine/http-client/internal/httpmethods"
	"github.com/alexflint/go-arg"
)

func main() {
	input := httpmethods.Input{}
	arg.MustParse(&input)

	httpResponse, err := httpmethods.RunHttpMethod(input)
	if err != nil {
		log.Fatal(err)
	}

	// Handle the HTTP response based on the command
	switch input.HTTPMethod {
	case "httphead":
		// For HEAD request, only print the response status
		printColoredHeaders(httpResponse.Header)
	default:
		// For other requests, print the response body
		defer httpResponse.Body.Close()
		body, err := io.ReadAll(httpResponse.Body)
		if err != nil {
			log.Fatal(err)
		}
		printColoredBody(body)
		err = writeToFile(input.Output, body)
		if err != nil {
			log.Fatal(err)
		}
	}
}
