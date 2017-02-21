package categoryDao

import (
	"github.com/NiciiA/GoGraphQL/dataaccess/mongoAccess"
	"github.com/NiciiA/GoGraphQL/domain/model/categoryModel"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var collectionName = "category"

var session *mgo.Session

var CategoryList map[bson.ObjectId]categoryModel.Category = make(map[bson.ObjectId]categoryModel.Category)

func GetCollection() *mgo.Collection {
	return mongoAccess.GetCollection(session, collectionName)
}

func GetByKey(key bson.ObjectId) categoryModel.Category {
	return CategoryList[key]
}

func AddCategory(c categoryModel.Category) {
	CategoryList[c.ID] = c
	Insert(c)
}

func UpdateCategory(c categoryModel.Category) {
	CategoryList[c.ID] = c
	Update(c)
}

func init() {
	session = mongoAccess.Session.Clone()
	var catList []categoryModel.Category
	GetAll().All(&catList)
	for _, cat := range catList {
		AddCategory(cat)
	}
}

func Insert(category categoryModel.Category) {
	GetCollection().Insert(&category)
}

func Update(category categoryModel.Category) {
	GetCollection().Update(bson.M{"_id": category.ID}, &category)
}

func Delete(category categoryModel.Category) {
	delete(CategoryList, category.ID)
	GetCollection().Remove(category.ID)
}

func GetAll() *mgo.Query {
	return GetCollection().Find(bson.M{})
}