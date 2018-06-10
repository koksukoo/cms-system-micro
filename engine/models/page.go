package models

import (
	"time"

	"github.com/globalsign/mgo/bson"
)

// Page corresponds to normal webpage. Can also be a portion of webpage
type Page struct {
	ID        bson.ObjectId     `bson:"_id" json:"id"`
	Slug      string            `bson:"slug" json:"slug"`
	Title     string            `bson:"title" json:"title"`
	Template  string            `bson:"template" json:"template"`
	Contents  map[string]string `bson:"contents" json:"contents"`
	Created   time.Time         `bson:"created" json:"created"`
	Modified  time.Time         `bson:"modified" json:"modified"`
	IsActive  bool              `bson:"is_active" json:"isActive"`
	ProjectID string            `bson:"project_id" json:"projectId"`
	Ancestors []string          `bson:"ancestors" json:"ancestors"`
	Parent    string            `bson:"parent" json:"parent"`
}
