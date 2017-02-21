package newsDao

import (
	"gopkg.in/mgo.v2"
	"github.com/NiciiA/GoGraphQL/dataaccess/mongoAccess"
	"gopkg.in/mgo.v2/bson"
	"github.com/NiciiA/GoGraphQL/domain/model/newsModel"
)

var collectionName = "news"

var session *mgo.Session

func GetCollection() *mgo.Collection {
	return mongoAccess.GetCollection(session, collectionName)
}

func GetById(id bson.ObjectId) *mgo.Query {
	return GetCollection().FindId(id)
}

func Insert(news newsModel.News) {
	GetCollection().Insert(&news)
}

func Update(news newsModel.News) {
	GetCollection().Update(bson.M{"_id": news.ID}, &news)
}

func Delete(news newsModel.News) {
	GetCollection().Remove(news.ID)
}

func GetAll(query interface{}) *mgo.Query {
	return GetCollection().Find(query)
}
func init() {
	session = mongoAccess.Session.Clone()
}