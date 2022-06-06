package api

import (
	"net/http"

	"cornerstone_verifier/pkg/config"
)

func NewRouter(config *config.Config) *http.ServeMux {
	r := http.NewServeMux()

	apiBaseURL := config.GetAPIBaseURL()

	// health
	r.HandleFunc(apiBaseURL+"/cornerstone/issuer/health", health(config))

	return r
}
