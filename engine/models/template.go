package models

import (
	"time"

	"github.com/globalsign/mgo/bson"
)

// PageTemplate contains mustache-style template
type PageTemplate struct {
	ID        bson.ObjectId `bson:"_id" json:"-"`
	ProjectID string        `bson:"project_id" json:"projectId"`
	Slug      string        `bson:"slug" json:"slug"`
	Title     string        `bson:"title" json:"title"`
	Content   string        `bson:"content" json:"content"`
	Created   time.Time     `bson:"created" json:"created"`
	Modified  time.Time     `bson:"modified" json:"modified"`
	IsActive  bool          `bson:"is_active" json:"isActive"`
}
