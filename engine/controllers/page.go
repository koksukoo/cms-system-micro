package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/globalsign/mgo/bson"
	"github.com/gorilla/mux"
	"github.com/mikkokokkoniemi/cms-system-micro/engine/models"
)

// GetPagesHierarchy returns a json map of page id's and titles in correct hierarchy
func GetPagesHierarchy(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Not implemented")
}

// CreatePage creates new page and sets ancestor-field in place
func CreatePage(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var page models.Page
	if err := json.NewDecoder(r.Body).Decode(&page); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	page.ID = bson.NewObjectId()
	page.Slug, _ = hashid.Encode([]int{int(page.ID.Counter())})
	page.Created = time.Now()
	page.Modified = time.Now()

	if page.Parent == "" {
		page.Parent = page.ID.Hex()
	}

	parent, err := dao.FindPageByID(page.Parent)
	if err == nil {
		page.Ancestors = append(parent.Ancestors, page.Parent)
	}

	if err := dao.InsertPage(page); err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusCreated, page)
}

// GetPage returns a single page
func GetPage(w http.ResponseWriter, r *http.Request) {
	page, err := dao.FindPageByID(mux.Vars(r)["pageId"])
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, page)
}

// UpdatePage updates a single page
func UpdatePage(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var page models.Page
	if err := json.NewDecoder(r.Body).Decode(&page); err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	page.Modified = time.Now()
	parent, err := dao.FindPageByID(page.Parent)
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	page.Ancestors = append(parent.Ancestors, page.Parent)

	if err := dao.UpdatePage(page); err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, page)
}

// DeletePage deletes a sigle page
func DeletePage(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	if err := dao.DeletePage(mux.Vars(r)["pageId"]); err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, map[string]string{"message": "success"})
}
