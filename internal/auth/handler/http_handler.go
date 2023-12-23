package handler

import (
	"encoding/json"
	"errors"
	"github.com/ikbarfp/bumder/internal/auth"
	"github.com/ikbarfp/bumder/pkg/response"
	"net/http"
)

type HttpHandler struct {
	authService auth.Service
}

// NewHttp . . .
func NewHttp(authSvc auth.Service) *HttpHandler {
	return &HttpHandler{authService: authSvc}
}

// Login . . .
func (h *HttpHandler) Login(w http.ResponseWriter, req *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	res := response.HttpResponse{}

	var body auth.RequestLogin
	if err := json.NewDecoder(req.Body).Decode(&body); err != nil {
		res.Message = "failed to decode request"
		res.Errors = err.Error()
		byteRes, _ := json.Marshal(res)

		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write(byteRes)

		return
	}

	userAuth, err := h.authService.Login(req.Context(), body)
	if err != nil {
		res.Message = err.Error()
		res.Errors = "bad request"
		byteRes, _ := json.Marshal(res)

		switch {
		case errors.Is(err, response.ErrUserNotFound):
			w.WriteHeader(http.StatusNotFound)
		case errors.Is(err, response.ErrInvalidPin):
			w.WriteHeader(http.StatusBadRequest)
		default:
			w.WriteHeader(http.StatusInternalServerError)
		}
		_, _ = w.Write(byteRes)

		return
	}

	res.Message = "login succeeded"
	res.Data = userAuth
	byteRes, _ := json.Marshal(res)
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(byteRes)

	return
}
