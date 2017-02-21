package entityActivityDao

import (
	"gopkg.in/mgo.v2"
	"github.com/NiciiA/GoGraphQL/dataaccess/mongoAccess"
	"github.com/NiciiA/GoGraphQL/domain/model/activityModel"
	"gopkg.in/mgo.v2/bson"
)

var collectionName = "entityActivity"

var session *mgo.Session

func GetCollection() *mgo.Collection {
	return mongoAccess.GetCollection(session, collectionName)
}

func GetById(id bson.ObjectId) *mgo.Query {
	return GetCollection().FindId(id)
}

func Insert(activity activityModel.Activity) {
	GetCollection().Insert(&activity)
}

func Update(activity activityModel.Activity) {
	GetCollection().Update(activity.ID, &activity)
}

func Delete(activity activityModel.Activity) {
	GetCollection().Remove(activity.ID)
}

func GetAll(entityID string) *mgo.Query {
	return GetCollection().Find(bson.M{"referenceId": entityID, "referenceClass":"Entity"})
}

func init() {
	session = mongoAccess.Session.Clone()
}