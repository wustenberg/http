package middleware

import (
	"net/http"
)

// Middleware sits between http.Handlers.
// The Middleware needs to call ServeHTTP on the given handler for the request to proceed in the chain.
type Middleware func(http.Handler) http.Handler

// Add Middleware to the handler given in the first parameter.
// Middlewares are called from right to left in the parameter list.
func Add(h http.Handler, middlewares ...Middleware) http.Handler {
	for _, middleware := range middlewares {
		h = middleware(h)
	}
	return h
}
