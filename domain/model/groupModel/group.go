package groupModel

import "gopkg.in/mgo.v2/bson"

type Group  struct {
	ID bson.ObjectId `bson:"_id" json:"_id"`
	Disabled bool `bson:"disabled" json:"disabled"`
	CreatedDate string `bson:"createdDate" json:"createdDate"`
	ModifiedDate string `bson:"modifiedDate" json:"modifiedDate"`
	Name string `bson:"name" json:"name"`
}