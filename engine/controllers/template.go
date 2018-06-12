package controllers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/globalsign/mgo/bson"
	"github.com/gorilla/mux"
	"github.com/mikkokokkoniemi/cms-system-micro/engine/models"
)

// ListTemplates lists templates for user
func ListTemplates(w http.ResponseWriter, r *http.Request) {
	templates, err := dao.FindAllTemplates()
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, templates)
}

// GetTemplate returns a json response of page template
func GetTemplate(w http.ResponseWriter, r *http.Request) {
	template, err := dao.FindTemplateByField("slug", mux.Vars(r)["templateId"])
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, template)
}

// CreateTemplate creates a page template
func CreateTemplate(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var template models.PageTemplate
	if err := json.NewDecoder(r.Body).Decode(&template); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	template.ID = bson.NewObjectId()
	template.Slug, _ = hashid.Encode([]int{int(template.ID.Counter())})
	template.Created = time.Now()
	template.Modified = time.Now()
	if err := dao.InsertTemplate(template); err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusCreated, template)
}

// UpdateTemplate updates template
func UpdateTemplate(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var template models.PageTemplate
	if err := json.NewDecoder(r.Body).Decode(&template); err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	template.Modified = time.Now()
	if err := dao.UpdateTemplate(template); err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, template)
}

func DeleteTemplate(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	if err := dao.DeleteTemplateByField("slug", mux.Vars(r)["templateId"]); err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, map[string]string{"message": "success"})
}
