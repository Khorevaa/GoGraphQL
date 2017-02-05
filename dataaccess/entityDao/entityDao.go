package entityDao

import (
	"github.com/NiciiA/GoGraphQL/dataaccess/mongoAccess"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"github.com/NiciiA/GoGraphQL/domain/model/entityModel"
)

var collectionName = "entity"

func GetCollection() *mgo.Collection {
	return mongoAccess.GetCollection(collectionName)
}

func GetById(id bson.ObjectId) *mgo.Query {
	return GetCollection().FindId(id)
}

func Insert(entity entityModel.Entity) {
	GetCollection().Insert(&entity)
}

func Update(entity entityModel.Entity) {
	GetCollection().Update(entity.ID, &entity)
}

func Delete(entity entityModel.Entity) {
	GetCollection().Remove(entity.ID)
}

func GetAll() *mgo.Query {
	return GetCollection().Find(bson.M{})
}