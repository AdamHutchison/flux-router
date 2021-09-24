package router

import (
	m "github.com/AdamHutchison/flux-router/middleware"
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
	router := &Router{mux: mux.NewRouter()}
	applyDefaultMiddleware(router)

	return router
}

func applyDefaultMiddleware(r *Router) {
	r.mux.Use(m.LoggingMiddleware)
}
