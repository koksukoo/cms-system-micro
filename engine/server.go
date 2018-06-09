package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mikkokokkoniemi/cms-system-micro/engine/controllers"
	"github.com/mikkokokkoniemi/cms-system-micro/engine/database"
)

var dao = database.MongoDAO{}

func init() {
	dao.Server = "localhost"
	dao.Database = "cms_engine"
}

func main() {
	r := mux.NewRouter()
	pr := r.PathPrefix("/projects").Subrouter()
	pr.HandleFunc("/", controllers.CreateProject).Methods("POST")
	pr.HandleFunc("/{id}", controllers.GetProject).Methods("GET")
	pr.HandleFunc("/{id}", controllers.UpdateProject).Methods("PUT")
	pr.HandleFunc("/{id}", controllers.DeleteProject).Methods("DELETE")

	if err := http.ListenAndServe(":5000", r); err != nil {
		log.Fatal(err)
	}
}
