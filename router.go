package router

import (
	"net/http"

	"github.com/gorilla/mux"
)

type RouterInterface interface {
	GetMux() http.Handler
}

type Router struct {
	mux http.Handler
}

func (r *Router) CreateMux() {
	r.mux = mux.NewRouter()
} 

func (r *Router) GetMux() http.Handler {
	return r.mux
} 