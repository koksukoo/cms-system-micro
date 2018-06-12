package controllers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/globalsign/mgo/bson"
	"github.com/gorilla/mux"
	"github.com/mikkokokkoniemi/cms-system-micro/engine/models"
)

type hierarchyItem struct {
	Title    string
	Slug     string
	Children map[string]hierarchyItem
}

// GetPagesHierarchy returns a json map of page id's and titles in correct hierarchy
func GetPagesHierarchy(w http.ResponseWriter, r *http.Request) {
	pages, err := dao.FindAllPages()
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	// var hierarchy map[string]hierarchyItem
	var hierarchy = make(map[string]hierarchyItem)

	for _, page := range pages {
		hierarchy = appendToHierarchy(
			hierarchy,
			page.Ancestors,
			hierarchyItem{page.Title, page.Slug, make(map[string]hierarchyItem)})
	}

	respondJSON(w, http.StatusOK, hierarchy)
}

// map is safe to pass by value because it's alway a pointer
func appendToHierarchy(hierarchy map[string]hierarchyItem, ancestors []string, item hierarchyItem) map[string]hierarchyItem {

	if len(ancestors) > 0 {
		if ha, ok := hierarchy[ancestors[0]]; ok {
			ha.Children = appendToHierarchy(ha.Children, ancestors[1:], item)
		} else {
			hierarchy[ancestors[0]] = hierarchyItem{"", "", map[string]hierarchyItem{item.Slug: item}}
		}
		return hierarchy
	}

	var children map[string]hierarchyItem

	if parent, ok := hierarchy[item.Slug]; ok {
		children = parent.Children
	} else {
		children = make(map[string]hierarchyItem)
	}
	hierarchy[item.Slug] = hierarchyItem{item.Title, item.Slug, children}
	return hierarchy
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

	if page.Parent != "" {
		parent, err := dao.FindPageByField("slug", page.Parent)
		if err == nil {
			page.Ancestors = append(parent.Ancestors, page.Parent)
		}
	}

	if err := dao.InsertPage(page); err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusCreated, page)
}

// GetPage returns a single page
func GetPage(w http.ResponseWriter, r *http.Request) {
	page, err := dao.FindPageByField("slug", mux.Vars(r)["pageId"])
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
	parent, err := dao.FindPageByField("slug", page.Parent)
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
	if err := dao.DeletePageByField("slug", mux.Vars(r)["pageId"]); err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, map[string]string{"message": "success"})
}
