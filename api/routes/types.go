package routes

import (
	"net/http"

	"github.com/ggualbertosouza/game/api/middleware"
)

type HandlerFunc func(w http.ResponseWriter, r *http.Request) error

type Route struct {
	Method string
	Path string
	Handler HandlerFunc
	Middlewares []middleware.Middleware
}

type Registerable interface {
	Routes() []Route
}