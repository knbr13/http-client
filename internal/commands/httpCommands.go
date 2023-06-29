package commands

import (
	"net/http"

	"github.com/abdullah-alaadine/http-client/internal/httpmethods"
)

type httpCommand struct {
	name        string
	description string
	run         func([]string) (*http.Response, error)
}

func GetHttpCommands() map[string]httpCommand {
	return map[string]httpCommand{
		"httpget": {
			name:        "httpget",
			description: "Sends HTTP GET request",
			run:         httpmethods.RunGet,
		},
		"httppost": {
			name:        "httppost",
			description: "Sends HTTP POST request",
			run:         httpmethods.RunPost,
		},
		"httpput": {
			name:        "httpput",
			description: "Sends HTTP PUT request",
			run:         httpmethods.RunPut,
		},
		"httpdelete": {
			name:        "httpdelete",
			description: "Sends HTTP DELETE request",
			run:         httpmethods.RunDelete,
		},
		"httppatch": {
			name:        "httppatch",
			description: "Sends HTTP PATCH request",
			run:         httpmethods.RunPatch,
		},
		"httpoptions": {
			name:        "httpoptions",
			description: "Sends HTTP OPTIONS request",
			run:         httpmethods.RunOptions,
		},
		"httphead": {
			name:        "httphead",
			description: "Sends HTTP HEAD request",
			run:         httpmethods.RunHead,
		},
	}
}
