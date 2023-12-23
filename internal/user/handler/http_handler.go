package handler

import (
	"encoding/json"
	"github.com/ikbarfp/bumder/internal/user"
	"github.com/ikbarfp/bumder/pkg/response"
	"net/http"
)

// HttpHandler . . .
type HttpHandler struct {
	userService user.Service
}

// NewHttp . . .
func NewHttp(userSvc user.Service) *HttpHandler {
	return &HttpHandler{userService: userSvc}
}

// Register . . .
func (h HttpHandler) Register(w http.ResponseWriter, req *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	res := response.HttpResponse{}

	var body user.RequestRegister
	if err := json.NewDecoder(req.Body).Decode(&body); err != nil {
		res.Message = "failed to decode request"
		res.Errors = err.Error()
		byteRes, _ := json.Marshal(res)

		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write(byteRes)

		return
	}

	if err := h.userService.Register(req.Context(), body); err != nil {
		res.Message = "bad request"
		res.Errors = err.Error()
		byteRes, _ := json.Marshal(res)

		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write(byteRes)

		return
	}

	res.Message = "success register new user"
	byteRes, _ := json.Marshal(res)

	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(byteRes)

	return
}

// Profile . . .
func (h HttpHandler) Profile(w http.ResponseWriter, req *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	res := response.HttpResponse{
		Message: "success get profile",
		Data: &user.User{
			Name:      "Kylian Mbappe",
			IsPremium: true,
		},
	}

	byteRes, _ := json.Marshal(res)
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(byteRes)
}
