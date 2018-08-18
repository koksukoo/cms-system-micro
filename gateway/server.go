package main

import (
	"log"
	"net/http"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

type Config struct {
	Server   confServer
	Projects confProjects
}

type confServer struct {
	Port string
}

type confProjects struct {
	Endpoint string
}

var conf Config

func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// if not authenticated return err

		// if authenticated call next handler
		next.ServeHTTP(w, r)
	})
}

// HandleProjectsRequest redirects request to engine service
func HandleProjectsRequest(w http.ResponseWriter, r *http.Request) {
	url := conf.Projects.Endpoint
	engineProxy := NewProxy(url)
	engineProxy.handle(w, r)
}

func main() {
	r := mux.NewRouter()
	pr := r.PathPrefix("/projects").Subrouter()
	pr.HandleFunc("/", HandleProjectsRequest)
	pr.Use(authMiddleware)

	loggedRouter := handlers.LoggingHandler(os.Stdout, r)
	log.Fatal(http.ListenAndServe(conf.Server.Port, loggedRouter))
}

func init() {
	if _, err := toml.DecodeFile("./config.toml", &conf); err != nil {
		log.Fatal(err)
	}
}
