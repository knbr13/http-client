package httpmethods

import (
	"bytes"
	"fmt"
	"net/http"

	"github.com/abdullah-alaadine/http-client/internal/utilities"
)

// RunPatch executes the PATCH request using the provided input parameters.
func RunPatch(input []string) (*http.Response, error) {
	url, body, headers, err := parsePatchInput(input)
	if err != nil {
		return nil, err
	}

	return patch(url, body, headers)
}

// Patch sends an HTTP PATCH request.
func patch(url string, body []byte, headers map[string]string) (*http.Response, error) {
	// Create an HTTP request with the request body
	httpRequest, err := http.NewRequest(http.MethodPatch, url, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}

	// Set headers
	for key, value := range headers {
		httpRequest.Header.Set(key, value)
	}

	httpRequest.Header.Set("Content-Type", "application/json")

	// Send the HTTP request
	httpResponse, err := httpClient.Do(httpRequest)
	if err != nil {
		return nil, err
	}

	return httpResponse, nil
}

// parsePatchInput parses the input slice for the PATCH command and extracts the URL, body, and headers.
func parsePatchInput(input []string) (string, []byte, map[string]string, error) {
	if len(input) < 2 || len(input) > 3 {
		return "", nil, nil, fmt.Errorf("invalid input for PATCH command")
	}

	url := input[0]
	bodyStr := input[1]
	headersStr := ""
	if len(input) == 3 {
		headersStr = input[2]
	}

	body, err := utilities.ParseBody(bodyStr)
	if err != nil {
		return "", nil, nil, err
	}

	headers, err := utilities.ParseHeaders(headersStr)
	if err != nil {
		return "", nil, nil, err
	}

	return url, body, headers, nil
}
