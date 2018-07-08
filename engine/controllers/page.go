package controllers

import (
	"encoding/json"
	"net/http"
	"sort"
	"time"

	"github.com/globalsign/mgo/bson"
	"github.com/gorilla/mux"
	"github.com/mikkokokkoniemi/cms-system-micro/engine/models"
)

/*
	Internally page controller models hierarchy in map, but outside it is returned as
	slice, because it makes more sense that way.
*/

type hierarchySlice []slicedHierarchyItem
type hierarchyMap map[string]hierarchyItem

type hierarchyItem struct {
	Title    string
	Slug     string
	Position int
	Children hierarchyMap
	Created  time.Time
	Template string
	IsActive bool
}

type slicedHierarchyItem struct {
	Title    string         `json:"title"`
	Slug     string         `json:"slug"`
	Position int            `json:"position"`
	Children hierarchySlice `json:"children"`
	Created  time.Time      `json:"created"`
	Template string         `json:"template"`
	IsActive bool           `json:"isActive"`
}

// hierarchySlice implements sort.Sort interface to slicedHierarchyItem
func (h hierarchySlice) Len() int      { return len(h) }
func (h hierarchySlice) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h hierarchySlice) Less(i, j int) bool {
	if h[i].Position == h[j].Position {
		return h[i].Created.Before(h[j].Created)
	}
	return h[i].Position < h[j].Position
}

// SortedSlice transforms map to a slice sorted by page position
func (h hierarchyMap) SortedSlice() []slicedHierarchyItem {
	var sorted []slicedHierarchyItem

	for _, item := range h {
		var sliced slicedHierarchyItem

		if len(item.Children) > 0 {
			sliced = slicedHierarchyItem{item.Title, item.Slug, item.Position, item.Children.SortedSlice(), item.Created, item.Template, item.IsActive}
		} else {
			sliced = slicedHierarchyItem{item.Title, item.Slug, item.Position, []slicedHierarchyItem{}, item.Created, item.Template, item.IsActive}
		}
		sorted = append(sorted, sliced)
	}
	sort.Sort(hierarchySlice(sorted))
	return sorted
}

// GetPagesHierarchy returns a json map of page id's and titles in correct hierarchy
func GetPagesHierarchy(w http.ResponseWriter, r *http.Request) {
	pages, err := dao.FindAllPages()
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	if len(pages) == 0 {
		respondJSON(w, http.StatusOK, []string{})
		return
	}

	// var hierarchy map[string]hierarchyItem
	var hierarchy = make(hierarchyMap)

	for _, page := range pages {
		hierarchy = appendToHierarchy(
			hierarchy,
			page.Ancestors,
			hierarchyItem{page.Title, page.Slug, page.Position, make(map[string]hierarchyItem), page.Created, page.Template, page.IsActive})
	}
	// hierarchy = appendToHierarchy(hierarchy, []string{}, hierarchyItem{"", "", 0, hierarchyMap{}, time.Now(), "", false})

	respondJSON(w, http.StatusOK, hierarchy.SortedSlice())
}

// map is safe to pass by value because it's alway a pointer
func appendToHierarchy(hierarchy map[string]hierarchyItem, ancestors []string, item hierarchyItem) map[string]hierarchyItem {

	if len(ancestors) > 0 {
		if ha, ok := hierarchy[ancestors[0]]; ok {
			ha.Children = appendToHierarchy(ha.Children, ancestors[1:], item)
		} else {
			hierarchy[ancestors[0]] = hierarchyItem{"", "", 0, map[string]hierarchyItem{item.Slug: item}, time.Now(), "", false}
		}
		return hierarchy
	}

	var children map[string]hierarchyItem

	if parent, ok := hierarchy[item.Slug]; ok {
		children = parent.Children
	} else {
		children = make(map[string]hierarchyItem)
	}
	hierarchy[item.Slug] = hierarchyItem{item.Title, item.Slug, item.Position, children, item.Created, item.Template, item.IsActive}

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
