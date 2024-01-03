package main

import (
	"io"
	"log"
	"strings"

	"github.com/alexflint/go-arg"

	"github.com/knbr13/http-client/internal/httpmethods"
)

func main() {
	input := httpmethods.Input{}
	arg.MustParse(&input)

	httpResponse, err := httpmethods.RunHttpMethod(input)
	if err != nil {
		log.Fatal(err)
	}
	defer httpResponse.Body.Close()

	switch strings.ToLower(input.HTTPMethod) {
	case "head":
		// For HEAD request, only print the response status
		printColoredHeaders(httpResponse.Header)
	default:
		// For other requests, print the response body
		body, err := io.ReadAll(httpResponse.Body)
		if err != nil {
			log.Fatal(err)
		}
		printColoredBody(body)
		if input.Output != "" {
			err = writeToFile(input.Output, body)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}
