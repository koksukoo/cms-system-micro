package models

import (
	"time"

	"github.com/globalsign/mgo/bson"
)

// Page corresponds to normal webpage. Can also be a portion of webpage
type Page struct {
	ID        bson.ObjectId   `bson:"_id" json:"id"`
	Title     string          `bson:"title" json:"title"`
	Template  bson.ObjectId   `bson:"template" json:"template"`
	Contents  []string        `bson:"contents" json:"contents"`
	Created   time.Time       `bson:"created" json:"created"`
	Modified  time.Time       `bson:"modified" json:"modified"`
	IsActive  bool            `bson:"is_active" json:"isActive"`
	ProjectID bson.ObjectId   `bson:"project_id" json:"projectId"`
	Ancestors []bson.ObjectId `bson:"ancestors" json:"ancestors"`
	Parent    bson.ObjectId   `bson:"parent" json:"parent"`
}
