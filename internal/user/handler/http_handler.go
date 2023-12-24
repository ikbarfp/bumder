package handler

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/ikbarfp/bumder/internal/user"
	"github.com/ikbarfp/bumder/pkg/response"
	"net/http"
)

// HttpHandler . . .
type HttpHandler struct {
	validator   *validator.Validate
	userService user.Service
}

// NewHttp . . .
func NewHttp(valid *validator.Validate, userSvc user.Service) *HttpHandler {
	return &HttpHandler{validator: valid, userService: userSvc}
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

	if err := h.validator.Struct(body); err != nil {
		res.Message = "validation errors"
		res.Errors = err.Error()
		byteRes, _ := json.Marshal(res)

		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write(byteRes)

		return
	}

	if err := h.userService.Register(req.Context(), body); err != nil {
		res.Message = err.Error()
		res.Errors = "bad request"
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
	res := response.HttpResponse{}

	pathParams := mux.Vars(req)
	userID := pathParams["user_id"]

	usr, err := h.userService.GetProfile(req.Context(), userID)
	if err != nil {
		res.Message = err.Error()
		res.Errors = "profile not found"
		byteRes, _ := json.Marshal(res)

		w.WriteHeader(http.StatusNotFound)
		_, _ = w.Write(byteRes)

		return
	}

	res.Message = "success get profile"
	res.Data = usr
	byteRes, _ := json.Marshal(res)
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(byteRes)

	return
}
