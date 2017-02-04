package groupModel

import "gopkg.in/mgo.v2/bson"

type Group  struct {
	ID bson.ObjectId `bson:"_id" json:"_id"`
	Name string `bson:"name" json:"name"`
}