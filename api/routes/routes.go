package routes

import (
	"net/http"

	"github.com/ggualbertosouza/game/api/handlers"
	"github.com/ggualbertosouza/game/api/middleware"

	"github.com/gorilla/mux"
)

type AppRoutes struct{}

func NewAppRoutes() *AppRoutes {
	return &AppRoutes{}
}

func (r *AppRoutes) Routes() []Route {
	return []Route{
		{
			Method:  http.MethodGet,
			Path:    "/health",
			Handler: handlers.HealthHandler,
		},
	}
}

func RegisterRoutes(router *mux.Router, registerable Registerable) {
	for _, route := range registerable.Routes() {
		handler := middleware.ChainMidds(
			middleware.ErrorMidd(adaptHandler(route.Handler)),
			route.Middlewares...,
		)
		router.HandleFunc(route.Path, handler).Methods(route.Method)
	}
}

func adaptHandler(h HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := h(w, r); err != nil {
			panic(err)
		}
	}
}