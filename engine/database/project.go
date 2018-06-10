package database

import (
	"github.com/globalsign/mgo/bson"
	"github.com/mikkokokkoniemi/cms-system-micro/engine/models"
)

// FindAllProjects finds all projects for current user.
// TODO: change to fetch only current user projects
func (m *MongoDAO) FindAllProjects() ([]models.Project, error) {
	var projects []models.Project
	err := db.C(ProjectCollection).Find(bson.M{}).All(&projects)
	return projects, err
}

// FindProjectByID fetches a single project entry
func (m *MongoDAO) FindProjectByID(id string) (models.Project, error) {
	var project models.Project
	err := db.C(ProjectCollection).FindId(bson.ObjectIdHex(id)).One(&project)
	return project, err
}

// InsertProject creates a project entity in db
func (m *MongoDAO) InsertProject(project models.Project) error {
	err := db.C(ProjectCollection).Insert(&project)
	return err
}

// UpdateProject updates a project entity in db
func (m *MongoDAO) UpdateProject(project models.Project) error {
	err := db.C(ProjectCollection).UpdateId(project.ID, &project)
	return err
}

// DeleteProject deletes a project by id
func (m *MongoDAO) DeleteProject(id string) error {
	err := db.C(ProjectCollection).RemoveId(bson.ObjectIdHex(id))
	return err
}
