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
	Insert(c)
}

func UpdateCategory(c categoryModel.Category) {
	CategoryList[c.Name] = c
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

func Insert(catgory categoryModel.Category) {
	GetCollection().Insert(&catgory)
}

func Update(catgory categoryModel.Category) {
	GetCollection().Update(catgory.ID, &catgory)
}

func Delete(catgory categoryModel.Category) {
	delete(CategoryList, catgory.Name)
	GetCollection().Remove(catgory.ID)
}

func GetAll() *mgo.Query {
	return GetCollection().Find(bson.M{})
}