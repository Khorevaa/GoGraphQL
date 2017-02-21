package contactDao

import (
	"gopkg.in/mgo.v2"
	"github.com/NiciiA/GoGraphQL/dataaccess/mongoAccess"
	"gopkg.in/mgo.v2/bson"
	"github.com/NiciiA/GoGraphQL/domain/model/contactModel"
)

var collectionName = "contact"

var session *mgo.Session

func GetCollection() *mgo.Collection {
	return mongoAccess.GetCollection(session, collectionName)
}

func GetById(id bson.ObjectId) *mgo.Query {
	return GetCollection().FindId(id)
}

func Insert(contact contactModel.Contact) {
	GetCollection().Insert(&contact)
}

func Update(contact contactModel.Contact) {
	GetCollection().Update(bson.M{"_id": contact.ID}, &contact)
}

func Delete(contact contactModel.Contact) {
	GetCollection().Remove(contact.ID)
}

func GetAll(query interface{}) *mgo.Query {
	return GetCollection().Find(query)
}
func init() {
	session = mongoAccess.Session.Clone()
}