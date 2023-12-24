package app

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
	_authHandler "github.com/ikbarfp/bumder/internal/auth/handler"
	_authRepo "github.com/ikbarfp/bumder/internal/auth/repository"
	_authService "github.com/ikbarfp/bumder/internal/auth/service"
	_feedsHandler "github.com/ikbarfp/bumder/internal/feeds/handler"
	_feedsService "github.com/ikbarfp/bumder/internal/feeds/service"
	_userHandler "github.com/ikbarfp/bumder/internal/user/handler"
	_userRepo "github.com/ikbarfp/bumder/internal/user/repository"
	_userService "github.com/ikbarfp/bumder/internal/user/service"
	"github.com/ikbarfp/bumder/pkg/response"
	"net/http"
)

func (s *Server) initValidator() {
	s.validator = validator.New()
}

func (s *Server) initRoutes() {
	s.router.Path("/").HandlerFunc(s.healthcheckHandler)
	s.router.Methods(http.MethodGet).Path("/healthcheck").HandlerFunc(s.healthcheckHandler)

	apiV1 := s.router.PathPrefix("/api/v1").Subrouter()

	authRepo := _authRepo.New()
	authService := _authService.New(authRepo)
	authHandler := _authHandler.NewHttp(authService)

	authV1 := apiV1.PathPrefix("/auth").Subrouter()
	authV1.Methods(http.MethodPost).Path("/login").HandlerFunc(authHandler.Login)

	userRepo := _userRepo.New()
	userService := _userService.New(userRepo, authRepo)
	userHandler := _userHandler.NewHttp(s.validator, userService)

	userV1 := apiV1.PathPrefix("/user").Subrouter()
	userV1.Methods(http.MethodPost).Path("/register").HandlerFunc(userHandler.Register)
	userV1.Methods(http.MethodGet).Path("/profile/{user_id}").HandlerFunc(userHandler.Profile)

	feedsService := _feedsService.New(userRepo)
	feedsHandler := _feedsHandler.NewHttp(feedsService)

	feedsV1 := apiV1.PathPrefix("/feeds").Subrouter()
	feedsV1.Methods(http.MethodGet).Path("/unseen/{user_id}").HandlerFunc(feedsHandler.Unseen)
	feedsV1.Methods(http.MethodPost).Path("/like").HandlerFunc(feedsHandler.Like)
	feedsV1.Methods(http.MethodPost).Path("/pass").HandlerFunc(feedsHandler.Pass)
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
