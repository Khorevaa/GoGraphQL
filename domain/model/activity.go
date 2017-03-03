package model

import "gopkg.in/mgo.v2/bson"

type Activity  struct {
	ID bson.ObjectId `bson:"_id" json:"_id"`
	Disabled bool `bson:"disabled" json:"disabled"`
	CreatedDate string `bson:"createdDate" json:"createdDate"`
	ModifiedDate string `bson:"modifiedDate" json:"modifiedDate"`
	ReferenceClass string `bson:"referenceClass" json:"referenceClass"`
	ReferenceId string `bson:"referenceId" json:"referenceId"`
	Comment string `bson:"comment" json:"comment"`
	Intern bool `bson:"intern" json:"intern"`
	Creator string `bson:"creator" json:"creator"`
}