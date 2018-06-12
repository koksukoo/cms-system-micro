package database

import (
	"github.com/globalsign/mgo/bson"
	"github.com/mikkokokkoniemi/cms-system-micro/engine/models"
)

// FindAllPages finds all projects for current user.
// TODO: change to fetch only current user pages
func (m *MongoDAO) FindAllPages() ([]models.Page, error) {
	var pages []models.Page
	err := db.C(PageCollection).Find(bson.M{}).All(&pages)
	return pages, err
}

// FindPageByID fetches a single page entry
func (m *MongoDAO) FindPageByID(id string) (models.Page, error) {
	var page models.Page
	err := db.C(PageCollection).FindId(bson.ObjectIdHex(id)).One(&page)
	return page, err
}

// FindPageByField fetches a single page entry by given field
func (m *MongoDAO) FindPageByField(field string, value string) (models.Page, error) {
	var page models.Page
	err := db.C(PageCollection).Find(bson.M{field: value}).One(&page)
	return page, err
}

// InsertPage creates a page entity in db
func (m *MongoDAO) InsertPage(page models.Page) error {
	err := db.C(PageCollection).Insert(&page)
	return err
}

// UpdatePage updates a page entity in db
func (m *MongoDAO) UpdatePage(page models.Page) error {
	err := db.C(PageCollection).UpdateId(page.ID, &page)
	return err
}

// DeletePage deletes a page by id
func (m *MongoDAO) DeletePage(id string) error {
	err := db.C(PageCollection).RemoveId(bson.ObjectIdHex(id))
	return err
}

// DeletePageByField deletes a page by given field
func (m *MongoDAO) DeletePageByField(field string, value string) error {
	err := db.C(PageCollection).Remove(bson.M{field: value})
	return err
}
