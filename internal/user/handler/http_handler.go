package handler

import (
	"encoding/json"
	"github.com/ikbarfp/bumder/internal/user"
	"github.com/ikbarfp/bumder/pkg/response"
	"net/http"
)

// HttpHandler . . .
type HttpHandler struct {
}

// NewHttpHandler . . .
func NewHttpHandler() *HttpHandler {
	return &HttpHandler{}
}

// Register . . .
func (h *HttpHandler) Register(w http.ResponseWriter, req *http.Request) {

	res := response.HttpResponse{
		Message: "success sign up",
	}

	byteRes, _ := json.Marshal(res)
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(byteRes)
}

// Profile . . .
func (h *HttpHandler) Profile(w http.ResponseWriter, req *http.Request) {

	res := response.HttpResponse{
		Data: &user.User{
			Name:        "Kylian Mbappe",
			MobilePhone: "+6286969696969",
			IsPremium:   true,
		},
	}

	byteRes, _ := json.Marshal(res)
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(byteRes)
}
