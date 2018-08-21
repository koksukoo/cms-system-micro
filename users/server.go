package main

import (
	"database/sql"
	"fmt"
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

	r.HandleFunc("/users/login", loginController).Methods("POST")
	r.HandleFunc("/users/register", registerController).Methods("POST")

	corsHandler := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: false,
		AllowedMethods:   []string{"GET", "POST", "OPTIONS", "PUT", "DELETE"},
		ExposedHeaders:   []string{"Access-Control-Allow-Origin", "Date", "Server", "Keep-Alive", "Connection", "Transfer-Encoding", "Content-Type"},
	}).Handler(r)

	loggedRouter := handlers.LoggingHandler(os.Stdout, corsHandler)
	log.Fatal(http.ListenAndServe(conf.Server.Port, loggedRouter))
}

func initDB() {
	connString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		conf.Database.User,
		conf.Database.Password,
		conf.Database.Server,
		conf.Database.Port,
		conf.Database.Database)
	conn, err := sql.Open("mysql", connString)
	if err != nil {
		log.Fatal(err.Error())
	}
	db = conn
}

func initCache() {
	connString := fmt.Sprintf("redis://%s%s", conf.Redis.Server, conf.Redis.Port)
	conn, err := redis.DialURL(connString)
	if err != nil {
		log.Fatal(err.Error())
	}
	cache = conn
}
