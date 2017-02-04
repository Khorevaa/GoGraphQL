package tagModel

import "gopkg.in/mgo.v2/bson"

type Tag  struct {
	ID bson.ObjectId `bson:"_id" json:"_id"`
	Name string `bson:"name" json:"name"`
	Style string `bson:"stye" json:"stye"`
}