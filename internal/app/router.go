package app

import (
	"encoding/json"
	_authHandler "github.com/ikbarfp/bumder/internal/auth/handler"
	_feedsHandler "github.com/ikbarfp/bumder/internal/feeds/handler"
	_userHandler "github.com/ikbarfp/bumder/internal/user/handler"
	"github.com/ikbarfp/bumder/pkg/response"
	"net/http"
)

func (s *Server) initRoutes() {
	s.router.Path("/").HandlerFunc(s.healthcheckHandler)
	s.router.Methods(http.MethodGet).Path("/healthcheck").HandlerFunc(s.healthcheckHandler)

	apiV1 := s.router.PathPrefix("/api/v1").Subrouter()

	authV1 := apiV1.PathPrefix("/auth").Subrouter()
	authHandler := _authHandler.NewHttpHandler()
	authV1.Methods(http.MethodPost).Path("/login").HandlerFunc(authHandler.Login)

	userV1 := apiV1.PathPrefix("/user").Subrouter()
	userHandler := _userHandler.NewHttpHandler()
	userV1.Methods(http.MethodPost).Path("/register").HandlerFunc(userHandler.Register)
	userV1.Methods(http.MethodGet).Path("/profile").HandlerFunc(userHandler.Profile)

	feedsV1 := apiV1.PathPrefix("/feeds").Subrouter()
	feedsHandler := _feedsHandler.NewHttpHandler()
	feedsV1.Methods(http.MethodGet).Path("/unseen").HandlerFunc(feedsHandler.Unseen)
	feedsV1.Methods(http.MethodPost).Path("/like").HandlerFunc(feedsHandler.Like)
	feedsV1.Methods(http.MethodPost).Path("/pass").HandlerFunc(feedsHandler.Pass)
	feedsV1.Methods(http.MethodPost).Path("/premium").HandlerFunc(feedsHandler.Premium)
}

func (s *Server) healthcheckHandler(w http.ResponseWriter, req *http.Request) {

	res := response.HttpResponse{
		Message: "OK",
	}

	byteRes, _ := json.Marshal(res)
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(byteRes)
}
