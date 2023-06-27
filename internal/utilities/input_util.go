package utilities

import (
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
