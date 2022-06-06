package server

import (
	"net/http"

	"cornerstone_verifier/pkg/log"
)

// NewLogRequest returns middleware request logger.
func NewLogRequest(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Info.Printf("%s - %s %s %s", r.RemoteAddr, r.Proto, r.Method, r.URL.RequestURI())
		next(w, r)
	}
}

// Middleware returns middleware functions.
type Middleware func(http.HandlerFunc) http.HandlerFunc

// ChainMiddleware chains all middleware functions right to left.
func ChainMiddleware(f http.HandlerFunc, m ...Middleware) http.HandlerFunc {
	// if our chain is done, use the original handlerfunc
	if len(m) == 0 {
		return f
	}

	// otherwise run recursively over nested handlers
	return m[0](ChainMiddleware(f, m[1:]...))
}
