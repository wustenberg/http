package middleware

import (
	"context"
	"net/http"
	"time"
)

// Timeout requests after a certain duration, by using context cancels.
func Timeout(after time.Duration) Middleware {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			ctx, cancel := context.WithTimeout(ctx, after)
			defer cancel()
			h.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
