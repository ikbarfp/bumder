package middleware

import (
	"github.com/google/uuid"
	"net/http"
)

// RequestID . . .
func RequestID(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		if req.Header.Get("X-Request-Id") == "" {
			req.Header.Set("X-Request-Id", uuid.New().String())
		}

		h.ServeHTTP(w, req.WithContext(req.Context()))

	})
}
