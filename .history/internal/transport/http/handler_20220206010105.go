package http

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Handler struct {
	Route mux.Router
}

func NewRouter() *Handler {
	return &Handler{}
}

func (h *Handler) SetupRoutes() {
	fmt.Println("starting routers")
	h.Route = mux.NewRouter()
	h.Route.HandleFunc("/api/mor", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hello Mor, I m up \n")
	})
}
