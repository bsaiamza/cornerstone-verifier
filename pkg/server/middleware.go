package server

import (
	"net/http"

	"iamza_verifier/pkg/log"
)

type Middleware func(http.HandlerFunc) http.HandlerFunc

func ChainMiddleware(f http.HandlerFunc, m ...Middleware) http.HandlerFunc {
	// if our chain is done, use the original HandlerFunc
	if len(m) == 0 {
		return f
	}

	// otherwise run recursively over nested handlers
	return m[0](ChainMiddleware(f, m[1:]...))
}

func LogAPIRequest(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Info.Printf("%s - %s %s %s", r.RemoteAddr, r.Proto, r.Method, r.URL.RequestURI())
		next(w, r)
	}
}
