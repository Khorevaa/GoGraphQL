package entityModel

import "gopkg.in/mgo.v2/bson"

type Entity  struct {
	ID bson.ObjectId `bson:"_id" json:"_id"`
	Disabled bool `bson:"disabled" json:"disabled"`
	CreatedDate string `bson:"createdDate" json:"createdDate"`
	ModifiedDate string `bson:"modifiedDate" json:"modifiedDate"`
	Subject string `bson:"subject" json:"subject"`
	Description string `bson:"description" json:"description"`
	Tags []string `bson:"tags" json:"tags"`
	Groups []string `bson:"groups" json:"groups"`
	Priority string `bson:"priority" json:"priority"`
	Category string `bson:"category" json:"category"`
	Creator string `bson:"creator" json:"creator"`
	Longitude string `bson:"longitude" json:"longitude"`
	Latitude string `bson:"latitude" json:"latitude"`
}