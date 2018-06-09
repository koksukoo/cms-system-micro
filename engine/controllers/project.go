package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/globalsign/mgo/bson"

	"github.com/mikkokokkoniemi/cms-system-micro/engine/database"
	"github.com/mikkokokkoniemi/cms-system-micro/engine/models"
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

// CreateProject creates a new project
func CreateProject(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var project models.Project
	if err := json.NewDecoder(r.Body).Decode(&project); err != nil {
		respondError(w, http.StatusBadRequest, "Bad request, baby")
		return
	}
	project.ID = bson.NewObjectId()
	if err := dao.InsertProject(project); err != nil {
		respondError(w, http.StatusInternalServerError, "Oh no! It doesn't work")
		return
	}
	respondJSON(w, http.StatusOK, project)
}

// GetProject returns a project hierarchy
func GetProject(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "not implemented yet!")
}

// UpdateProject updates project hierarchy and owners
func UpdateProject(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "not implemented yet!")
}

// DeleteProject removes project completely
func DeleteProject(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "not implemented yet!")
}
