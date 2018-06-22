package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mikkokokkoniemi/cms-system-micro/engine/controllers"
	"github.com/rs/cors"
)

func main() {
	// allowedHeaders := handlers.AllowedHeaders([]string{"X-Requested-With"})
	// allowedOrigins := handlers.AllowedOrigins([]string{"*"})
	// allowedMethods := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "DELETE", "OPTIONS"})

	r := mux.NewRouter()
	// https://hostname/projects/
	pr := r.PathPrefix("/projects").Subrouter()
	pr.HandleFunc("/", controllers.ListProjects).Methods("GET")
	pr.HandleFunc("/", controllers.CreateProject).Methods("POST")
	pr.HandleFunc("/{projectId}", controllers.GetProject).Methods("GET")
	pr.HandleFunc("/{projectId}", controllers.UpdateProject).Methods("PUT")
	pr.HandleFunc("/{projectId}", controllers.DeleteProject).Methods("DELETE")
	// https://hostname/projects/{id}/pages
	qr := pr.PathPrefix("/{projectId}/pages").Subrouter()
	qr.HandleFunc("/", controllers.GetPagesHierarchy).Methods("GET")
	qr.HandleFunc("/", controllers.CreatePage).Methods("POST")
	qr.HandleFunc("/{pageId}", controllers.GetPage).Methods("GET")
	qr.HandleFunc("/{pageId}", controllers.UpdatePage).Methods("PUT")
	qr.HandleFunc("/{pageId}", controllers.DeletePage).Methods("DELETE")
	// https://hostname/projects/{id}/templates
	tr := pr.PathPrefix("/{projectId}/templates").Subrouter()
	tr.HandleFunc("/", controllers.ListTemplates).Methods("GET")
	tr.HandleFunc("/", controllers.CreateTemplate).Methods("POST")
	tr.HandleFunc("/{templateId}", controllers.GetTemplate).Methods("GET")
	tr.HandleFunc("/{templateId}", controllers.UpdateTemplate).Methods("PUT")
	tr.HandleFunc("/{templateId}", controllers.DeleteTemplate).Methods("DELETE")

	handler := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
		AllowedMethods:   []string{"GET", "POST", "OPTIONS", "PUT", "DELETE"},
		ExposedHeaders:   []string{"Access-Control-Allow-Origin", "Date", "Server", "Keep-Alive", "Connection", "Transfer-Encoding", "Content-Type"},
	}).Handler(r)

	log.Fatal(http.ListenAndServe(":6000", handler))

}
