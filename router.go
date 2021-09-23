package router

import (
	"github.com/gorilla/mux"
)

type RouterInterface interface {
	GetMux() *mux.Router
}

type Router struct {
	mux *mux.Router
}

func (r *Router) GetMux() *mux.Router {
	return r.mux
}

func NewRouter() *Router {
	router := Router{mux: mux.NewRouter()}

	return &router
}
