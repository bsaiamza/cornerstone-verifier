package api

import (
	"net/http"

	"cornerstone_verifier/pkg/config"
	"cornerstone_verifier/pkg/server"
)

func health(config *config.Config) http.HandlerFunc {
	mdw := []server.Middleware{
		server.NewLogRequest,
	}

	return server.ChainMiddleware(healthHandler(config), mdw...)
}
func healthHandler(config *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	}
}
