package handler

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/ikbarfp/bumder/internal/feeds"
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

// Unseen . . .
func (h *HttpHandler) Unseen(w http.ResponseWriter, req *http.Request) {

	res := response.HttpResponse{
		Data: []*feeds.UnseenUser{
			{
				UserID:    uuid.NewString(),
				Name:      "Erling Haaland",
				IsPremium: false,
				Age:       22,
			},
			{
				UserID:    uuid.NewString(),
				Name:      "Jude Bellingham",
				IsPremium: true,
				Age:       20,
			},
			{
				UserID:    uuid.NewString(),
				Name:      "Vinicius Junior",
				IsPremium: false,
				Age:       21,
			},
		},
	}

	byteRes, _ := json.Marshal(res)
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(byteRes)
}

// Like . . .
func (h *HttpHandler) Like(w http.ResponseWriter, req *http.Request) {

	res := response.HttpResponse{
		Message: "success to liked user",
	}

	byteRes, _ := json.Marshal(res)
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(byteRes)
}

// Pass . . .
func (h *HttpHandler) Pass(w http.ResponseWriter, req *http.Request) {

	res := response.HttpResponse{
		Message: "success to pass user",
	}

	byteRes, _ := json.Marshal(res)
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(byteRes)
}

// Premium . . .
func (h *HttpHandler) Premium(w http.ResponseWriter, req *http.Request) {

	res := response.HttpResponse{
		Message: "success to buy premium",
	}

	byteRes, _ := json.Marshal(res)
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(byteRes)
}
