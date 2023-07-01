package httpmethods

import (
	"net/http"
	"time"
)

var httpClient *http.Client = &http.Client{
	Timeout: 30 * time.Second,
}
