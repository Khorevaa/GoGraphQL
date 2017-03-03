package categoryDao

import (
	"github.com/NiciiA/GoGraphQL/dataaccess/mongoAccess"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"github.com/NiciiA/GoGraphQL/domain/model"
)

var collectionName = "category"

var session *mgo.Session

var CategoryList map[bson.ObjectId]model.Category = make(map[bson.ObjectId]model.Category)

func GetCollection() *mgo.Collection {
	return mongoAccess.GetCollection(session, collectionName)
}

func GetByKey(key bson.ObjectId) model.Category {
	return CategoryList[key]
}

func AddCategory(c model.Category) {
	CategoryList[c.ID] = c
	Insert(c)
}

func UpdateCategory(c model.Category) {
	CategoryList[c.ID] = c
	Update(c)
}

func init() {
	session = mongoAccess.Session.Clone()
	var catList []model.Category
	GetAll().All(&catList)
	for _, cat := range catList {
		AddCategory(cat)
	}
}

func Insert(category model.Category) {
	GetCollection().Insert(&category)
}

func Update(category model.Category) {
	GetCollection().Update(bson.M{"_id": category.ID}, &category)
}

func Delete(category model.Category) {
	delete(CategoryList, category.ID)
	GetCollection().Remove(category.ID)
}

func GetAll() *mgo.Query {
	return GetCollection().Find(bson.M{})
}

func SearchAll(search string) *mgo.Query {
	return GetCollection().Find(bson.M{"$text": bson.M{"$search": search}})
}