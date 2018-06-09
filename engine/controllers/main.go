package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/mikkokokkoniemi/cms-system-micro/engine/database"
)

var dao = database.MongoDAO{}

func init() {
	dao.Server = "localhost"
	dao.Database = "cms_engine"
	dao.Connect()
}

func respondError(w http.ResponseWriter, code int, msg string) {
	respondJSON(w, code, map[string]string{"error": msg})
}

func respondJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
