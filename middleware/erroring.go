package middleware

import (
	"math/rand"
	"net/http"
)

// Erroring injects HTTP 500 errors with a certain probability.
// Probability should be between 0 and 1, both inclusive.
func Erroring(probability float64) Middleware {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if rand.Float64() <= probability {
				http.Error(w, "You've been selected for a random error by middleware", http.StatusInternalServerError)
			}
			h.ServeHTTP(w, r)
		})
	}
}

