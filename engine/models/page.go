package models

import (
	"time"

	"github.com/globalsign/mgo/bson"
)

// Page corresponds to normal webpage. Can also be a portion of webpage
type Page struct {
	ID        bson.ObjectId     `bson:"_id" json:"-"`
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

/*
 Tässä tarvitaan nyt rekursiivista ajattelua

 parenttiin mätsäävät lapsiksi -> lapsiin mätsäävät lapsiksi -> lapsiin mätsäävät lapsiksi
 matchParent(hierarchyItem)

 makeAncestors([x, 3, 5])
 jos ansestoria ei ole, luodaan template sille, jos on, ei tehdä mitään
 ancestor{id: $id, chidlren: makeAncestor(ancestors[:1])}
*/
