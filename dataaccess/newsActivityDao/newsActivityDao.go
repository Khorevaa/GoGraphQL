package newsActivityDao

import (
	"gopkg.in/mgo.v2"
	"github.com/NiciiA/GoGraphQL/dataaccess/mongoAccess"
	"gopkg.in/mgo.v2/bson"
	"github.com/NiciiA/GoGraphQL/domain/model"
)

var collectionName = "newsActivity"

var session *mgo.Session

func GetCollection() *mgo.Collection {
	return mongoAccess.GetCollection(session, collectionName)
}

func GetById(id bson.ObjectId) *mgo.Query {
	return GetCollection().FindId(id)
}

func Insert(activity model.Activity) {
	GetCollection().Insert(&activity)
}

func Update(activity model.Activity) {
	GetCollection().Update(bson.M{"_id": activity.ID}, &activity)
}

func Delete(activity model.Activity) {
	GetCollection().Remove(activity.ID)
}

func GetAll(entityID string) *mgo.Query {
	return GetCollection().Find(bson.M{"referenceId": entityID, "referenceClass":"News"})
}

func init() {
	session = mongoAccess.Session.Clone()
}

func SearchAll(search string) *mgo.Query {
	return GetCollection().Find(bson.M{"$text": bson.M{"$search": search}})
}