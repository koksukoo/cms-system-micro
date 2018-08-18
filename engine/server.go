package main

import (
	"log"
	"net/http"
	"os"

	"github.com/BurntSushi/toml"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/mikkokokkoniemi/cms-system-micro/engine/controllers"
	"github.com/rs/cors"
)

type Config struct {
	Server confServer
}

type confServer struct {
	Port string
}

var conf Config

func init() {
	if _, err := toml.DecodeFile("./config.toml", &conf); err != nil {
		log.Fatal(err)
	}
}

func main() {
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

	corsHandler := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: false,
		AllowedMethods:   []string{"GET", "POST", "OPTIONS", "PUT", "DELETE"},
		ExposedHeaders:   []string{"Access-Control-Allow-Origin", "Date", "Server", "Keep-Alive", "Connection", "Transfer-Encoding", "Content-Type"},
	}).Handler(r)

	loggedRouter := handlers.LoggingHandler(os.Stdout, corsHandler)

	log.Fatal(http.ListenAndServe(conf.Server.Port, loggedRouter))

}
