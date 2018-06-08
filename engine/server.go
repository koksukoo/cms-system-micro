package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mikkokokkoniemi/cms-system-micro/engine/controllers"
)

func main() {
	r := mux.NewRouter()
	pr := r.PathPrefix("/projects").Subrouter()
	pr.HandleFunc("/", controllers.CreateProject).Methods("POST")
	pr.HandleFunc("/{id}", controllers.GetProject).Methods("GET")
	pr.HandleFunc("/{id}", controllers.UpdateProject).Methods("PUT")
	pr.HandleFunc("/{id}", controllers.DeleteProject).Methods("Delete")

	if err := http.ListenAndServe(":5000", r); err != nil {
		log.Fatal(err)
	}
}
