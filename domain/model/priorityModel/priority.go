package priorityModel

import "gopkg.in/mgo.v2/bson"

type Priority  struct {
	ID bson.ObjectId `bson:"_id" json:"_id"`
	Name string `bson:"name" json:"name"`
}