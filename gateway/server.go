package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// if not authenticated return err

		// if authenticated call next handler
		next.ServeHTTP(w, r)
	})
}

// HandleProjectsRequest redirects request to engine service
func HandleProjectsRequest(w http.ResponseWriter, r *http.Request) {
	url := "http://localhost:6000"
	engineProxy := NewProxy(url)
	engineProxy.handle(w, r)
}

func main() {
	r := mux.NewRouter()
	pr := r.PathPrefix("/projects").Subrouter()
	pr.HandleFunc("/", HandleProjectsRequest)
	pr.Use(authMiddleware)

	loggedRouter := handlers.LoggingHandler(os.Stdout, r)
	log.Fatal(http.ListenAndServe(":5004", loggedRouter))
}
