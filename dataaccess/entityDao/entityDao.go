package entityDao

import (
	"github.com/NiciiA/GoGraphQL/dataaccess/mongoAccess"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"github.com/NiciiA/GoGraphQL/domain/model"
)

var collectionName = "entity"

var session *mgo.Session

func GetCollection() *mgo.Collection {
	return mongoAccess.GetCollection(session, collectionName)
}

func GetById(id bson.ObjectId) *mgo.Query {
	return GetCollection().FindId(id)
}

func Insert(entity model.Entity) {
	GetCollection().Insert(&entity)
}

func Update(entity model.Entity) {
	GetCollection().Update(bson.M{"_id": entity.ID}, &entity)
}

func Delete(entity model.Entity) {
	GetCollection().Remove(entity.ID)
}

func GetAll(query interface{}) *mgo.Query {
	return GetCollection().Find(query)
}
func init() {
	session = mongoAccess.Session.Clone()
}

func SearchAll(search string) *mgo.Query {
	c := GetCollection()
	c.EnsureIndexKey("_id", "subject", "description")
	return c.Find(bson.M{"$text": bson.M{"$search": search}})
}