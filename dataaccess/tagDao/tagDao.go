package tagDao

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"github.com/NiciiA/GoGraphQL/domain/model"
	"github.com/NiciiA/GoGraphQL/dataaccess/mongoAccess"
)

var collectionName = "tag"

var session *mgo.Session

var TagList map[bson.ObjectId]model.Tag = make(map[bson.ObjectId]model.Tag)

func GetCollection() *mgo.Collection {
	return mongoAccess.GetCollection(session, collectionName)
}

func GetByKey(key bson.ObjectId) model.Tag {
	return TagList[key]
}

func AddTag(c model.Tag) {
	TagList[c.ID] = c
	Insert(c)
}

func UpdateTag(c model.Tag) {
	TagList[c.ID] = c
	Update(c)
}

func init() {
	session = mongoAccess.Session.Clone()
	var tagList []model.Tag
	GetAll().All(&tagList)
	for _, tag := range tagList {
		AddTag(tag)
	}
}

func Insert(tag model.Tag) {
	GetCollection().Insert(&tag)
}

func Update(tag model.Tag) {
	GetCollection().Update(bson.M{"_id": tag.ID}, &tag)
}

func Delete(tag model.Tag) {
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