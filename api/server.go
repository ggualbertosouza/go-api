package api

import (
	"context"
	"log"
	"net/http"

	"github.com/ggualbertosouza/game/api/middleware"
	"github.com/ggualbertosouza/game/api/routes"

	"github.com/gorilla/mux"
)

type AppServer struct {
	httpServer *http.Server
	router *mux.Router
}

func New(addr string) *AppServer {
	router := mux.NewRouter()

	router.Use(middleware.JSONContentType, middleware.LogMidd)

	routes.RegisterRoutes(router, routes.NewAppRoutes())

	return &AppServer{
		httpServer: &http.Server{
			Addr:    addr,
			Handler: router,
		},
		router: router,
	}
}

func (s *AppServer) Start() error {
	log.Printf("Server running on %s", s.httpServer.Addr)
	return s.httpServer.ListenAndServe()
}

func (s *AppServer) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}