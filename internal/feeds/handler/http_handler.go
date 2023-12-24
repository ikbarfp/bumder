package handler

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/ikbarfp/bumder/internal/feeds"
	"github.com/ikbarfp/bumder/pkg/response"
	"net/http"
)

// HttpHandler . . .
type HttpHandler struct {
	feedsService feeds.Service
}

// NewHttp . . .
func NewHttp(feedsSvc feeds.Service) *HttpHandler {
	return &HttpHandler{feedsService: feedsSvc}
}

// Unseen . . .
func (h HttpHandler) Unseen(w http.ResponseWriter, req *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	res := response.HttpResponse{}

	pathParams := mux.Vars(req)
	userID := pathParams["user_id"]

	unseenUsers, err := h.feedsService.Unseen(req.Context(), userID)
	if err != nil {
		res.Message = err.Error()
		res.Errors = "bad request"
		byteRes, _ := json.Marshal(res)

		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write(byteRes)

		return
	}

	res.Message = "success get data"
	res.Data = unseenUsers
	byteRes, _ := json.Marshal(res)
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(byteRes)

	return
}

// Like . . .
func (h HttpHandler) Like(w http.ResponseWriter, req *http.Request) {

	res := response.HttpResponse{
		Message: "success to liked user",
	}

	byteRes, _ := json.Marshal(res)
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(byteRes)
}

// Pass . . .
func (h HttpHandler) Pass(w http.ResponseWriter, req *http.Request) {

	res := response.HttpResponse{
		Message: "success to pass user",
	}

	byteRes, _ := json.Marshal(res)
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(byteRes)
}
