package http

import (
	"fmt"
	"gitlab.com/zenport.io/go-assignment/engine"
	"gitlab.com/zenport.io/go-assignment/handlers"
	"net/http"

	"github.com/gorilla/mux"
)

type HTTPAdapter struct{}

var r *mux.Router

func (adapter *HTTPAdapter) Start() {
	fmt.Println("Starting HTTP connection...")
	// TODO: Get port number from environment variable
	http.ListenAndServe(":8000", r)
}

func (adapter *HTTPAdapter) Stop() {
	r = nil
}

func NewHTTPAdapter(e engine.Engine) *HTTPAdapter {
	r = mux.NewRouter()

	r.HandleFunc("/knight", handlers.HandleFindAll(e).ServeHTTP).Methods("GET")
	r.HandleFunc("/knight/{id}", handlers.HandleFindOne(e).ServeHTTP).Methods("GET")
	r.HandleFunc("/knight", handlers.HandleCreate(e).ServeHTTP).Methods("POST")
	r.HandleFunc("/fight", handlers.HandleFight(e).ServeHTTP).Methods("POST")
	http.Handle("/", r)
	return &HTTPAdapter{}
}
