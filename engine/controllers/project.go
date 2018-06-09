package controllers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/mux"

	"github.com/globalsign/mgo/bson"

	"github.com/mikkokokkoniemi/cms-system-micro/engine/models"
)

// CreateProject creates a new project
func CreateProject(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var project models.Project
	if err := json.NewDecoder(r.Body).Decode(&project); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	project.ID = bson.NewObjectId()
	project.Created = time.Now()
	project.Modified = time.Now()
	if err := dao.InsertProject(project); err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusCreated, project)
}

// ListProjects lists projects for user
func ListProjects(w http.ResponseWriter, r *http.Request) {
	projects, err := dao.FindAllProjects()
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, projects)
}

// GetProject returns a project hierarchy
func GetProject(w http.ResponseWriter, r *http.Request) {
	project, err := dao.FindProjectByID(mux.Vars(r)["id"])
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, project)
}

// UpdateProject updates project hierarchy and owners
func UpdateProject(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var project models.Project
	if err := json.NewDecoder(r.Body).Decode(&project); err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	project.Modified = time.Now()
	if err := dao.UpdateProject(project); err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, project)
}

// DeleteProject removes project completely
func DeleteProject(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	if err := dao.DeleteProject(mux.Vars(r)["id"]); err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, map[string]string{"message": "success"})
}
