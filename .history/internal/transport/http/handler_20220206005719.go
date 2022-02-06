package http

import (
	"github.com/gorilla/mux"
)

type Handler struct {
	Route mux.Router
}

func NewRouter() *Handler {
	return &Handler{}
}
