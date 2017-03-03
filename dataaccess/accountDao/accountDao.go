package accountDao

import (
	"gopkg.in/mgo.v2"
	"github.com/NiciiA/GoGraphQL/dataaccess/mongoAccess"
	"gopkg.in/mgo.v2/bson"
	"github.com/NiciiA/GoGraphQL/domain/model"
)

var collectionName = "account"

var session *mgo.Session

func GetCollection() *mgo.Collection {
	return mongoAccess.GetCollection(session, collectionName)
}

func GetById(id bson.ObjectId) *mgo.Query {
	return GetCollection().FindId(id)
}

func Insert(account model.Account) {
	GetCollection().Insert(&account)
}

func Update(account model.Account) {
	GetCollection().Update(bson.M{"_id": account.ID}, &account)
}

func Delete(account model.Account) {
	GetCollection().Remove(account.ID)
}

func GetAll(query interface{}) *mgo.Query {
	return GetCollection().Find(query)
}
func init() {
	session = mongoAccess.Session.Clone()
}

func SearchAll(search string) *mgo.Query {
	return GetCollection().Find(bson.M{"$text": bson.M{"$search": search}})
}