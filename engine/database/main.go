package database

import (
	"log"

	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"github.com/mikkokokkoniemi/cms-system-micro/engine/models"
)

// MongoDAO is a database access object for mongo server
type MongoDAO struct {
	Server   string
	Database string
}

var db *mgo.Database

// Collection names
const (
	ProjectCollection  = "engine_projects"
	PageCollection     = "engine_pages"
	TemplateCollection = "engine_templates"
)

// Connect establishes connection to db
func (m *MongoDAO) Connect() {
	session, err := mgo.Dial(m.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(m.Database)
}

// Close ends the connection to db. This should not be used when application closes
func (m *MongoDAO) Close() {
	if db == nil {
		return
	}
	db.Session.Close()
}

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
