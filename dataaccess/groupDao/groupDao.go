package groupDao

import (
"github.com/NiciiA/GoGraphQL/dataaccess/mongoAccess"
"github.com/NiciiA/GoGraphQL/domain/model/groupModel"
"gopkg.in/mgo.v2"
"gopkg.in/mgo.v2/bson"
)

var collectionName = "group"

var session *mgo.Session

var GroupList map[string]groupModel.Group = make(map[string]groupModel.Group)

func GetCollection() *mgo.Collection {
	return mongoAccess.GetCollection(session, collectionName)
}

func GetByKey(key string) groupModel.Group {
	return GroupList[key]
}

func AddGroup(c groupModel.Group) {
	GroupList[c.Name] = c
	Insert(c)
}

func UpdateGroup(c groupModel.Group) {
	GroupList[c.Name] = c
	Update(c)
}

func init() {
	session = mongoAccess.Session.Clone()
	var groupList []groupModel.Group
	GetAll().All(&groupList)
	for _, grp := range groupList {
		AddGroup(grp)
	}
}

func Insert(group groupModel.Group) {
	GetCollection().Insert(&group)
}

func Update(group groupModel.Group) {
	GetCollection().Update(group.ID, &group)
}

func Delete(group groupModel.Group) {
	delete(GroupList, group.Name)
	GetCollection().Remove(group.ID)
}

func GetAll() *mgo.Query {
	return GetCollection().Find(bson.M{})
}