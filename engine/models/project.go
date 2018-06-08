package models

import (
	"time"

	"github.com/globalsign/mgo/bson"
)

// Project is the top-tier model in system
type Project struct {
	ID       bson.ObjectId `bson:"_id" json:"id"`
	Title    string        `bson:"title" json:"title"`
	Created  time.Time     `bson:"created" json:"created"`
	Modified time.Time     `bson:"modified" json:"modified"`
	Owners   []string      `bson:"owners" json:"owners"`
	IsActive bool          `bson:"is_active" json:"isActive"`
}
