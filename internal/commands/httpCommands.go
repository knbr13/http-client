package commands

import (
	"net/http"
	"strings"

	"github.com/abdullah-alaadine/http-client/internal/httpmethods"
)

type httpCommand struct {
	name        string
	description string
	run         func(string, []byte, map[string]string) (*http.Response, error)
}

func getHttpCommands() map[string]httpCommand {
	return map[string]httpCommand{
		"httpget": {
			name:        "httpget",
			description: "Sends HTTP GET request",
			run:         httpmethods.Get,
		},
		"httppost": {
			name:        "httppost",
			description: "Sends HTTP POST request",
			run:         httpmethods.Post,
		},
		"httpput": {
			name:        "httpput",
			description: "Sends HTTP PUT request",
			run:         httpmethods.Put,
		},
		"httpdelete": {
			name:        "httpdelete",
			description: "Sends HTTP DELETE request",
			run:         httpmethods.Delete,
		},
		"httppatch": {
			name:        "httppatch",
			description: "Sends HTTP PATCH request",
			run:         httpmethods.Patch,
		},
		"httpoptions": {
			name:        "httpoptions",
			description: "Sends HTTP OPTIONS request",
			run:         httpmethods.Options,
		},
		"httphead": {
			name:        "httphead",
			description: "Sends HTTP HEAD request",
			run:         httpmethods.Head,
		},
	}

}

func cleanInput(str string) []string {
    lowerCase := strings.ToLower(str)
    words := strings.Fields(lowerCase)
    return words
}