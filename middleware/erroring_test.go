package middleware_test

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/wustenberg/http/middleware"
)

func TestErroring(t *testing.T) {
	t.Run("sends http 500", func(t *testing.T) {
		h := middleware.Add(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}), middleware.Erroring(1))

		recorder := httptest.NewRecorder()
		h.ServeHTTP(recorder, httptest.NewRequest("", "/", nil))
		result := recorder.Result()
		if result.StatusCode != http.StatusInternalServerError {
			t.Fatalf("status code is %v, should be %v", result.StatusCode, http.StatusInternalServerError)
		}
	})

	t.Run("does not send errors if probability 0", func(t *testing.T) {
		h := middleware.Add(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			_, err := io.WriteString(w, "OK")
			if err != nil {
				panic(err)
			}
		}), middleware.Erroring(0))

		for i := 0; i < 10000; i++ {
			recorder := httptest.NewRecorder()
			h.ServeHTTP(recorder, httptest.NewRequest("", "/", nil))
			result := recorder.Result()
			if result.StatusCode != http.StatusOK {
				t.Fatalf("status code is %v, should be %v", result.StatusCode, http.StatusOK)
			}
		}
	})
}
