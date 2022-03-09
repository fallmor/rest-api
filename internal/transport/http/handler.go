package http

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/fallmor/rest-api/internal/comment"
	"github.com/fallmor/rest-api/internal/metrics"
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type Handler struct {
	Route   *mux.Router
	Service *comment.Service
}
type Response struct {
	Message string
	err     error
}

// NewHandler reoutourne un pointeur vers la structure Handler
func NewRouter(service *comment.Service) *Handler {
	return &Handler{
		Service: service,
	}
}

func init() {
	// Metrics have to be registered to be exposed:
	prometheus.MustRegister(metrics.Goroutine)
	prometheus.MustRegister(metrics.Countcpu)
}

// SetupRoutes ajoute les routes
func (h *Handler) SetupRoutes() {
	fmt.Println("starting routers")
	h.Route = mux.NewRouter()
	h.Route.HandleFunc("/api/mor/{id}", h.GetComment).Methods("GET")
	h.Route.HandleFunc("/api/mor", h.GetAllComment).Methods("GET")
	h.Route.HandleFunc("/api/mor", h.PostComment).Methods("POST")
	h.Route.HandleFunc("/api/mor/{id}", h.UpdateComment).Methods("PUT")
	h.Route.HandleFunc("/api/mor/{id}", h.DeleteComment).Methods("DELETE")

	// h.Route.Handle("/metrics", promhttp.HandlerFor(prometheus.DefaultGatherer, promhttp.HandlerOpts{
	// 	// Opt into OpenMetrics to support exemplars.
	// 	EnableOpenMetrics: true,
	// 	// need to use my custom metrics
	// },
	// ))
	h.Route.Handle("/metrics", promhttp.Handler())
	h.Route.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(Response{Message: "Hello mor, I'm up"}); err != nil {
			fmt.Fprintln(w, "Error performing request")
		}
	})
}

// GetComment - return a comment
func (h *Handler) GetComment(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	vars := mux.Vars(r)
	id := vars["id"]

	id_int, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		fmt.Fprintln(w, "Unable to parse Id to uint")
	}
	com, err := h.Service.GetComment(uint(id_int))
	if err != nil {
		fmt.Fprintln(w, "Error performing request")
	}
	if err := json.NewEncoder(w).Encode(com); err != nil {
		panic(err)
	}
	//fmt.Fprintln(w, "", jsonCom)

}

// var tmpl = template.Must(template.ParseFiles("forms.html"))
// use html templating

// PostComment - creates a new comment
func (h *Handler) PostComment(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	var comment comment.Comment

	if err := json.NewDecoder(r.Body).Decode(&comment); err != nil {
		fmt.Fprintln(w, "failed to decode the body")
	}
	comment, err := h.Service.PostComment(comment)
	if err != nil {
		fmt.Fprintln(w, "Error Posting the comment")
	}
	if err := json.NewEncoder(w).Encode(comment); err != nil {
		panic(err)
	}
	//fmt.Fprintln(w, "", comment)
}

// GetAllComments - return all comments
func (h *Handler) GetAllComment(w http.ResponseWriter, r *http.Request) {

	comments, err := h.Service.GetAllComments()
	if err != nil {
		fmt.Fprintln(w, "Error requestting all comments")
	}
	if err := json.NewEncoder(w).Encode(comments); err != nil {
		panic(err)
	}
	//fmt.Fprintln(w, "", jsonCom)
}

//  UpdateComment - update comment
func (h *Handler) UpdateComment(w http.ResponseWriter, r *http.Request) {
	var comment comment.Comment
	if err := json.NewDecoder(r.Body).Decode(&comment); err != nil {
		fmt.Fprintln(w, "failed to decode the body")
	}

	// Get the id

	vars := mux.Vars(r)
	id := vars["id"]

	id_int, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		fmt.Fprintln(w, "Unable to parse Id to uint")
	}
	Updatedcomment, err := h.Service.UpdateComment(uint(id_int), comment)
	if err != nil {
		fmt.Fprintln(w, "Error updating comment")
	}

	// jsonCom := json.NewEncoder(w).Encode(comment)
	// fmt.Fprintln(w, "", jsonCom)
	if err := json.NewEncoder(w).Encode(Updatedcomment); err != nil {
		panic(err)
	}
}

// DeleteComment -  delete comment
func (h *Handler) DeleteComment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	id_int, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		fmt.Fprintln(w, "Unable to parse Id to uint")
	}
	err = h.Service.DeleteComment(uint(id_int))
	responseErorr(w, err, "Error deleting comment")
	responseOk(w, "comment deleted successfully")
	// 	if err != nil {
	// 		fmt.Fprintln(w, "Error deleting comment")
	// 	}
	// 	fmt.Fprintln(w, "comment deleted successfully")
}

func responseErorr(w http.ResponseWriter, err error, message string) {
	w.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(w).Encode(Response{
		err:     err,
		Message: message,
	},
	)
}

func responseOk(w http.ResponseWriter, resp interface{}) {
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}
