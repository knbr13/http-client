package commands

import (
	"net/http"

	"github.com/abdullah-alaadine/http-client/internal/httpmethods"
)

type httpCommand struct {
	Name        string
	Description string
	Run         func([]string) (*http.Response, error)
}

func GetHttpCommands() map[string]httpCommand {
	return map[string]httpCommand{
		"httpget": {
			Name:        "httpget",
			Description: "Sends HTTP GET request",
			Run:         httpmethods.RunGet,
		},
		"httppost": {
			Name:        "httppost",
			Description: "Sends HTTP POST request",
			Run:         httpmethods.RunPost,
		},
		"httpput": {
			Name:        "httpput",
			Description: "Sends HTTP PUT request",
			Run:         httpmethods.RunPut,
		},
		"httpdelete": {
			Name:        "httpdelete",
			Description: "Sends HTTP DELETE request",
			Run:         httpmethods.RunDelete,
		},
		"httppatch": {
			Name:        "httppatch",
			Description: "Sends HTTP PATCH request",
			Run:         httpmethods.RunPatch,
		},
		"httpoptions": {
			Name:        "httpoptions",
			Description: "Sends HTTP OPTIONS request",
			Run:         httpmethods.RunOptions,
		},
		"httphead": {
			Name:        "httphead",
			Description: "Sends HTTP HEAD request",
			Run:         httpmethods.RunHead,
		},
	}
}
