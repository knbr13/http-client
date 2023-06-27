package commands

import (
	"net/http"

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
			description: "Sends an HTTP GET request",
			run:         httpmethods.Get,
		},
		"httppost": {
			name:        "httppost",
			description: "Sends an HTTP POST request",
			run:         httpmethods.Post,
		},
	}
}
