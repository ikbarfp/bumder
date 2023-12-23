package handler

import (
	"encoding/json"
	"github.com/ikbarfp/bumder/pkg/response"
	"net/http"
)

type HttpHandler struct {
}

func NewHttpHandler() *HttpHandler {
	return &HttpHandler{}
}

// Login . . .
func (h *HttpHandler) Login(w http.ResponseWriter, req *http.Request) {

	res := response.HttpResponse{
		Message: "login succeeded",
	}

	byteRes, _ := json.Marshal(res)
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(byteRes)
}
