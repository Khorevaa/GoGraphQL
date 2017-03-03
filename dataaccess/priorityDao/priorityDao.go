package priorityDao

import (
	"github.com/NiciiA/GoGraphQL/domain/model"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"github.com/NiciiA/GoGraphQL/dataaccess/mongoAccess"
)

var collectionName = "priority"

var session *mgo.Session

var PriorityList map[bson.ObjectId]model.Priority = make(map[bson.ObjectId]model.Priority)

func GetCollection() *mgo.Collection {
	return mongoAccess.GetCollection(session, collectionName)
}

func GetByKey(key bson.ObjectId) model.Priority {
	return PriorityList[key]
}

func AddPriority(c model.Priority) {
	PriorityList[c.ID] = c
	Insert(c)
}

func UpdatePriority(c model.Priority) {
	PriorityList[c.ID] = c
	Update(c)
}

func init() {
	session = mongoAccess.Session.Clone()
	var priorityList []model.Priority
	GetAll().All(&priorityList)
	for _, priority := range priorityList {
		AddPriority(priority)
	}
}

func Insert(priority model.Priority) {
	GetCollection().Insert(&priority)
}

func Update(priority model.Priority) {
	GetCollection().Update(bson.M{"_id": priority.ID}, &priority)
}

func Delete(priority model.Priority) {
	delete(PriorityList, priority.ID)
	GetCollection().Remove(priority.ID)
}

func GetAll() *mgo.Query {
	return GetCollection().Find(bson.M{})
}

func SearchAll(search string) *mgo.Query {
	return GetCollection().Find(bson.M{"$text": bson.M{"$search": search}})
}