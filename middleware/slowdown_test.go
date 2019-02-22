package middleware_test

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"http/middleware"
)

func TestSlowdown(t *testing.T) {
	t.Run("slows down handling", func(t *testing.T) {
		done := make(chan bool, 1)

		h := middleware.Add(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			done <- true
		}), middleware.Slowdown(time.Second))

		go func() {
			h.ServeHTTP(httptest.NewRecorder(), httptest.NewRequest("", "/", nil))
		}()

		select {
		case <-done:
			t.Fatal("received done too soon")
		case <-time.After(500 * time.Millisecond):
			// Do nothing
		}
	})
}
