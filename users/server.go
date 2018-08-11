package main

import (
	"database/sql"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

var db *sql.DB

func main() {
	initDB()
	r := mux.NewRouter()

	r.HandleFunc("/login", loginController).Methods("POST")
	r.HandleFunc("/register", registerController).Methods("POST")

	handler := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: false,
		AllowedMethods:   []string{"GET", "POST", "OPTIONS", "PUT", "DELETE"},
		ExposedHeaders:   []string{"Access-Control-Allow-Origin", "Date", "Server", "Keep-Alive", "Connection", "Transfer-Encoding", "Content-Type"},
	}).Handler(r)
	log.Fatal(http.ListenAndServe(":7000", handler))
}

func initDB() {
	conn, err := sql.Open("mysql", "root:secret@/cms_userdb")
	if err != nil {
		log.Fatal(err.Error())
	}
	db = conn
}
