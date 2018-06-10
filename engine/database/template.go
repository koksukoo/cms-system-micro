package database

import (
	"github.com/globalsign/mgo/bson"
	"github.com/mikkokokkoniemi/cms-system-micro/engine/models"
)

// FindAllTemplates finds all projects for current user.
// TODO: change to fetch only current user templates
func (m *MongoDAO) FindAllTemplates() ([]models.PageTemplate, error) {
	var templates []models.PageTemplate
	err := db.C(TemplateCollection).Find(bson.M{}).All(&templates)
	return templates, err
}

// FindTemplateByID fetches a single page entry
func (m *MongoDAO) FindTemplateByID(id string) (models.PageTemplate, error) {
	var template models.PageTemplate
	err := db.C(TemplateCollection).FindId(bson.ObjectIdHex(id)).One(&template)
	return template, err
}

// InsertTemplate creates a project entity in db
func (m *MongoDAO) InsertTemplate(template models.PageTemplate) error {
	err := db.C(TemplateCollection).Insert(&template)
	return err
}

// UpdateTemplate updates a project entity in db
func (m *MongoDAO) UpdateTemplate(template models.PageTemplate) error {
	err := db.C(TemplateCollection).UpdateId(template.ID, &template)
	return err
}

// DeleteTemplate deletes a project by id
func (m *MongoDAO) DeleteTemplate(id string) error {
	err := db.C(TemplateCollection).RemoveId(bson.ObjectIdHex(id))
	return err
}
