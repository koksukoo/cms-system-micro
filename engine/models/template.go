package models

import "github.com/globalsign/mgo/bson"

// PageTemplate contains mustache-style template
type PageTemplate struct {
	ID      bson.ObjectId `bson:"_id" json:"id"`
	Title   string        `bson:"title" json:"title"`
	Content string        `bson:"content" json:"content"`
}
