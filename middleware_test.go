package http_middleware_test

import (
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	middleware "http_middleware"
)

func TestAdd(t *testing.T) {
	t.Run("adds middleware called from right to left", func(t *testing.T) {
		var h http.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			_, err := io.WriteString(w, "hat")
			if err != nil {
				t.Fatal(err)
			}
		})

		calledFirst := func(h http.Handler) http.Handler {
			return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				_, err := io.WriteString(w, "po")
				if err != nil {
					t.Fatal(err)
				}
				h.ServeHTTP(w, r)
			})
		}

		calledSecond := func(h http.Handler) http.Handler {
			return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				_, err := io.WriteString(w, "ny")
				if err != nil {
					t.Fatal(err)
				}
				h.ServeHTTP(w, r)
			})
		}

		h = middleware.Add(h, calledSecond, calledFirst)

		recorder := httptest.NewRecorder()
		request := httptest.NewRequest("GET", "/", nil)
		h.ServeHTTP(recorder, request)

		result := recorder.Result()
		body, err := ioutil.ReadAll(result.Body)
		if err != nil {
			t.Fatal(err)
		}

		if string(body) != "ponyhat" {
			t.Fatal("got unexpected body", string(body))
		}
	})
}
