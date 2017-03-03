package model

import "gopkg.in/mgo.v2/bson"

type News  struct {
	ID bson.ObjectId `bson:"_id" json:"_id"`
	Disabled bool `bson:"disabled" json:"disabled"`
	CreatedDate string `bson:"createdDate" json:"createdDate"`
	ModifiedDate string `bson:"modifiedDate" json:"modifiedDate"`
	Title string `bson:"title" json:"title"`
	Text string `bson:"text" json:"text"`
	Intern bool `bson:"intern" json:"intern"`
	Tags []string `bson:"tags" json:"tags"`
	Groups []string `bson:"groups" json:"groups"`
	Important bool `bson:"important" json:"important"`
	Category string `bson:"category" json:"category"`
	Creator string `bson:"creator" json:"creator"`
}