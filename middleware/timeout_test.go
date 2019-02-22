package middleware_test

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/wustenberg/http/middleware"
)

func TestTimeout(t *testing.T) {
	t.Run("times out request after given duration", func(t *testing.T) {
		cancelled := make(chan bool, 1)

		h := middleware.Add(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			<-ctx.Done()
			cancelled <- true
		}), middleware.Timeout(100 * time.Millisecond))

		go func() {
			h.ServeHTTP(httptest.NewRecorder(), httptest.NewRequest("", "/", nil))
		}()

		select {
		case <-cancelled:
			// Do nothing
		case <-time.After(200 * time.Millisecond):
			t.Fatal("context not cancelled")
		}
	})
}
