package accountDao

import (
	"gopkg.in/mgo.v2"
	"github.com/NiciiA/GoGraphQL/dataaccess/mongoAccess"
	"gopkg.in/mgo.v2/bson"
	"github.com/NiciiA/GoGraphQL/domain/model/accountModel"
)

var collectionName = "account"

var session *mgo.Session

func GetCollection() *mgo.Collection {
	return mongoAccess.GetCollection(session, collectionName)
}

func GetById(id bson.ObjectId) *mgo.Query {
	return GetCollection().FindId(id)
}

func Insert(account accountModel.Account) {
	GetCollection().Insert(&account)
}

func Update(account accountModel.Account) {
	GetCollection().Update(account.ID, &account)
}

func Delete(account accountModel.Account) {
	GetCollection().Remove(account.ID)
}

func GetAll(query interface{}) *mgo.Query {
	return GetCollection().Find(query)
}
func init() {
	session = mongoAccess.Session.Clone()
}