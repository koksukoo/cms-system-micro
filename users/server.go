package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gomodule/redigo/redis"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

var db *sql.DB
var cache redis.Conn

func main() {
	initDB()
	initCache()
	r := mux.NewRouter()

	r.HandleFunc("/login", loginController).Methods("POST")
	r.HandleFunc("/register", registerController).Methods("POST")

	corsHandler := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: false,
		AllowedMethods:   []string{"GET", "POST", "OPTIONS", "PUT", "DELETE"},
		ExposedHeaders:   []string{"Access-Control-Allow-Origin", "Date", "Server", "Keep-Alive", "Connection", "Transfer-Encoding", "Content-Type"},
	}).Handler(r)

	loggedRouter := handlers.LoggingHandler(os.Stdout, corsHandler)
	log.Fatal(http.ListenAndServe(":7000", loggedRouter))
}

func initDB() {
	conn, err := sql.Open("mysql", "root:secret@/cms_userdb")
	if err != nil {
		log.Fatal(err.Error())
	}
	db = conn
}

func initCache() {
	conn, err := redis.DialURL("redis://localhost:6379")
	if err != nil {
		log.Fatal(err.Error())
	}
	cache = conn
}
