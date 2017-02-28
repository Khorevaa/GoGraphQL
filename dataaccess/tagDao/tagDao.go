package tagDao

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"github.com/NiciiA/GoGraphQL/domain/model/tagModel"
	"github.com/NiciiA/GoGraphQL/dataaccess/mongoAccess"
)

var collectionName = "tag"

var session *mgo.Session

var TagList map[bson.ObjectId]tagModel.Tag = make(map[bson.ObjectId]tagModel.Tag)

func GetCollection() *mgo.Collection {
	return mongoAccess.GetCollection(session, collectionName)
}

func GetByKey(key bson.ObjectId) tagModel.Tag {
	return TagList[key]
}

func AddTag(c tagModel.Tag) {
	TagList[c.ID] = c
	Insert(c)
}

func UpdateTag(c tagModel.Tag) {
	TagList[c.ID] = c
	Update(c)
}

func init() {
	session = mongoAccess.Session.Clone()
	var tagList []tagModel.Tag
	GetAll().All(&tagList)
	for _, tag := range tagList {
		AddTag(tag)
	}
}

func Insert(tag tagModel.Tag) {
	GetCollection().Insert(&tag)
}

func Update(tag tagModel.Tag) {
	GetCollection().Update(bson.M{"_id": tag.ID}, &tag)
}

func Delete(tag tagModel.Tag) {
	delete(TagList, tag.ID)
	GetCollection().Remove(tag.ID)
}

func GetAll() *mgo.Query {
	return GetCollection().Find(bson.M{})
}

func SearchAll(search string) *mgo.Query {
	c := GetCollection()
	c.EnsureIndexKey("_id", "name", "style")
	return c.Find(bson.M{"$regex": bson.M{"$search": search}})
}