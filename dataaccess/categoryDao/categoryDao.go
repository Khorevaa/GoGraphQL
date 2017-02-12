package categoryDao

import (
	"github.com/NiciiA/GoGraphQL/dataaccess/mongoAccess"
	"github.com/NiciiA/GoGraphQL/domain/model/categoryModel"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var collectionName = "category"

var session *mgo.Session

var CategoryList map[string]categoryModel.Category = make(map[string]categoryModel.Category)

func GetCollection() *mgo.Collection {
	return mongoAccess.GetCollection(session, collectionName)
}

func GetByKey(key string) categoryModel.Category {
	return CategoryList[key]
}

func AddCategory(c categoryModel.Category) {
	CategoryList[c.Name] = c
}

func init() {
	session = mongoAccess.Session.Clone()
	var catList []categoryModel.Category
	GetCollection().Find(bson.M{}).All(&catList)
	for _, cat := range catList {
		AddCategory(cat)
	}
}