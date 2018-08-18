package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/BurntSushi/toml"
	"github.com/mikkokokkoniemi/cms-system-micro/engine/database"
	hashids "github.com/speps/go-hashids"
)

type Config struct {
	Database confDB
	Hashids  confHash
}

type confDB struct {
	Server   string
	Database string
}

type confHash struct {
	Salt string
}

var dao = database.MongoDAO{}
var hashid = hashids.HashID{}
var conf Config

func init() {
	if _, err := toml.DecodeFile("./config.toml", &conf); err != nil {
		log.Fatal(err)
	}

	dao.Server = conf.Database.Server
	dao.Database = conf.Database.Database
	dao.Connect()

	hd := hashids.NewData()
	hd.Salt = conf.Hashids.Salt
	hd.MinLength = 8
	h, _ := hashids.NewWithData(hd)
	hashid = *h
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
