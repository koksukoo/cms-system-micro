package database

import (
	"log"

	"github.com/globalsign/mgo"
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

// Close ends the connection to db. Should be used when application closes
func (m *MongoDAO) Close() {
	if db == nil {
		return
	}
	db.Session.Close()
}
