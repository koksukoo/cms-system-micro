package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/gomodule/redigo/redis"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

type Config struct {
	Server   confServer
	Projects confEndpoint
	Users    confEndpoint
	Redis    confRedis
}

type confServer struct {
	Port string
}

type confEndpoint struct {
	Endpoint string
}

type confRedis struct {
	Server string
	Port   string
}

var conf Config
var cache redis.Conn

func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		sessionCookie, err := r.Cookie("cms_session")
		if err != nil {
			respondError(w, http.StatusUnauthorized, err.Error())
			return
		}
		cacheValue, err := cache.Do("GET", sessionCookie)
		if err != nil {
			respondError(w, http.StatusUnauthorized, err.Error())
			return
		}
		w.Header().Set("X-Cms-User", fmt.Sprintf("%v", cacheValue))
		next.ServeHTTP(w, r)
	})
}

func main() {
	r := mux.NewRouter()
	endpoint, _ := url.Parse(conf.Projects.Endpoint)
	ar := r.PathPrefix("/projects").Subrouter()
	ar.PathPrefix("/").Handler(httputil.NewSingleHostReverseProxy(endpoint))

	endpoint, _ = url.Parse(conf.Users.Endpoint)
	r.PathPrefix("/users").Handler(httputil.NewSingleHostReverseProxy(endpoint))
	ar.Use(authMiddleware)

	loggedRouter := handlers.LoggingHandler(os.Stdout, r)
	log.Fatal(http.ListenAndServe(conf.Server.Port, loggedRouter))
}

func init() {
	if _, err := toml.DecodeFile("./config.toml", &conf); err != nil {
		log.Fatal(err)
	}
	initCache()
}

func respondError(w http.ResponseWriter, code int, msg string) {
	respondJSON(w, code, map[string]string{"error": msg})
}

func respondJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.WriteHeader(code)
	w.Write(response)
}

func initCache() {
	connString := fmt.Sprintf("redis://%s%s", conf.Redis.Server, conf.Redis.Port)
	conn, err := redis.DialURL(connString)
	if err != nil {
		log.Fatal(err.Error())
	}
	cache = conn
}
