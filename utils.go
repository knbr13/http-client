package main

import (
	"fmt"
	"net/http"
	"strings"
)

func formatHeaders(headers http.Header) string {
	var formattedHeaders strings.Builder

	for key, values := range headers {
		formattedHeaders.WriteString(fmt.Sprintf("%s: %s\n", key, strings.Join(values, ", ")))
	}

	return formattedHeaders.String()
}
