package http_middleware

import (
	"net/http"
	"time"
)

// Slowdown the call to the next handler by the given duration.
func Slowdown(by time.Duration) Middleware {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			time.Sleep(by)
			h.ServeHTTP(w, r)
		})
	}
}
