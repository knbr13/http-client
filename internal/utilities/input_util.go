package utilities

import (
	"encoding/json"
	"fmt"
	"strings"
)

// ParseHeaders parses the input headers string into a map.
func ParseHeaders(headersStr string) (map[string]string, error) {
	headers := make(map[string]string)
	headersStr = strings.TrimSpace(headersStr)

	if len(headersStr) == 0 {
		return headers, nil
	}

	if !strings.HasPrefix(headersStr, "headers={") || !strings.HasSuffix(headersStr, "}") {
		return nil, fmt.Errorf("invalid headers format: %s", headersStr)
	}

	headersStr = strings.TrimPrefix(headersStr, "headers={")
	headersStr = strings.TrimSuffix(headersStr, "}")

	headerPairs := strings.Split(headersStr, ",")

	for _, pair := range headerPairs {
		parts := strings.Split(pair, ":")
		if len(parts) != 2 {
			return nil, fmt.Errorf("invalid header format: %s", pair)
		}
		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])
		headers[key] = value
	}

	return headers, nil
}

// ParseBody parses the input body string into a JSON byte slice.
func ParseBody(bodyStr string) ([]byte, error) {
	bodyStr = strings.TrimSpace(bodyStr)

	if len(bodyStr) == 0 {
		return nil, fmt.Errorf("body is required")
	}

	if !strings.HasPrefix(bodyStr, "body={") || !strings.HasSuffix(bodyStr, "}") {
		return nil, fmt.Errorf("invalid body format: %s", bodyStr)
	}

	bodyStr = strings.TrimPrefix(bodyStr, "body={")
	bodyStr = strings.TrimSuffix(bodyStr, "}")

	// Convert the body map to JSON bytes
	jsonBody, err := json.Marshal(bodyStr)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal body as JSON: %w", err)
	}

	return jsonBody, nil
}

func CleanInput(str string) []string {
	lowerCase := strings.ToLower(str)
	words := strings.Fields(lowerCase)
	return words
}
