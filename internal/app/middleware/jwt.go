package middleware

import (
	"encoding/json"
	"github.com/ikbarfp/bumder/pkg/response"
	"net/http"
)

// AuthToken . . .
func AuthToken(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		if req.Header.Get("Authorization") == "" {
			w.WriteHeader(http.StatusUnauthorized)

			res := response.HttpResponse{
				Message: "unauthorized request",
			}

			byteRes, _ := json.Marshal(res)
			_, _ = w.Write(byteRes)
		}
	})
}
