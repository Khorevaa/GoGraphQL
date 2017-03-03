package groupDao

import (
	"github.com/NiciiA/GoGraphQL/dataaccess/mongoAccess"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"github.com/NiciiA/GoGraphQL/domain/model"
)

var collectionName = "group"

var session *mgo.Session

var GroupList map[bson.ObjectId]model.Group = make(map[bson.ObjectId]model.Group)

func GetCollection() *mgo.Collection {
	return mongoAccess.GetCollection(session, collectionName)
}

func GetByKey(key bson.ObjectId) model.Group {
	return GroupList[key]
}

func AddGroup(c model.Group) {
	GroupList[c.ID] = c
	Insert(c)
}

func UpdateGroup(c model.Group) {
	GroupList[c.ID] = c
	Update(c)
}

func init() {
	session = mongoAccess.Session.Clone()
	var groupList []model.Group
	GetAll().All(&groupList)
	for _, grp := range groupList {
		AddGroup(grp)
	}
}

func Insert(group model.Group) {
	GetCollection().Insert(&group)
}

func Update(group model.Group) {
	GetCollection().Update(bson.M{"_id": group.ID}, &group)
}

func Delete(group model.Group) {
	delete(GroupList, group.ID)
	GetCollection().Remove(group.ID)
}

func GetAll() *mgo.Query {
	return GetCollection().Find(bson.M{})
}

func SearchAll(search string) *mgo.Query {
	return GetCollection().Find(bson.M{"$text": bson.M{"$search": search}})
}