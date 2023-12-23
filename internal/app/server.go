package app

import (
	"context"
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/ikbarfp/bumder/pkg/config"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

// Server . . .
type Server struct {
	router *mux.Router
}

// New . . .
func New() *Server {
	log.Println("--------------- Starting The Service ---------------")

	cfg := config.New(
		config.WithFilename("config"),
		config.WithExtension("yml"),
		config.WithPath("."),
	)

	if err := cfg.Load(); err != nil {
		log.Fatalf("[ERROR] load config : %v\n", err)
	}

	server := &Server{
		router: mux.NewRouter(),
	}
	server.initRoutes()

	return server
}

// Run . . .
func (s *Server) Run() error {
	s.printBanner()

	httpServer := &http.Server{
		Addr:    ":" + config.GetString("server_app.port"),
		Handler: s.router,
	}
	log.Println("listening on port :" + config.GetString("server_app.port"))

	go func() {
		if err := httpServer.ListenAndServe(); err != nil {
			if !errors.Is(err, http.ErrServerClosed) {
				log.Fatalf("[ERROR] failed to start server : %v", err)
			}
		}
	}()

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer func() {
		cancel()
	}()

	signCh := make(chan os.Signal, 1)
	signal.Notify(signCh, syscall.SIGTERM, syscall.SIGINT)

	systemCall := <-signCh
	log.Println(". . . received " + systemCall.String() + " signal")

	if err := httpServer.Shutdown(ctx); err != nil {
		log.Fatalf("[ERROR] failed to shutdown server : %v", err)
	}

	log.Println("--------------- Shutdown Service Completed ---------------")
	return nil
}

func (s *Server) printBanner() {
	banner := fmt.Sprintf("\n    __                        __         "+
		"\n   / /_  __  ______ ___  ____/ /__  _____"+
		"\n  / __ \\/ / / / __ `__ \\/ __  / _ \\/ ___/"+
		"\n / /_/ / /_/ / / / / / / /_/ /  __/ /    "+
		"\n/_.___/\\__,_/_/ /_/ /_/\\__,_/\\___/_/  %v   "+
		"\n                                         \n", config.GetString("server_app.version"))

	fmt.Print(banner)
}
