package orgUnitDao

import (
	"github.com/NiciiA/GoGraphQL/dataaccess/mongoAccess"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"github.com/NiciiA/GoGraphQL/domain/model"
)

var collectionName = "orgUnit"

var session *mgo.Session

var OrgUnitList map[bson.ObjectId]model.OrgUnit = make(map[bson.ObjectId]model.OrgUnit)

func GetCollection() *mgo.Collection {
	return mongoAccess.GetCollection(session, collectionName)
}

func GetByKey(key bson.ObjectId) model.OrgUnit {
	return OrgUnitList[key]
}

func AddOrgUnit(c model.OrgUnit) {
	OrgUnitList[c.ID] = c
	Insert(c)
}

func UpdateOrgUnit(c model.OrgUnit) {
	OrgUnitList[c.ID] = c
	Update(c)
}

func init() {
	session = mongoAccess.Session.Clone()
	var orgUnitList []model.OrgUnit
	GetAll().All(&orgUnitList)
	for _, orgU := range orgUnitList {
		AddOrgUnit(orgU)
	}
}

func Insert(orgUnit model.OrgUnit) {
	GetCollection().Insert(&orgUnit)
}

func Update(orgUnit model.OrgUnit) {
	GetCollection().Update(bson.M{"_id": orgUnit.ID}, &orgUnit)
}

func Delete(orgUnit model.OrgUnit) {
	delete(OrgUnitList, orgUnit.ID)
	GetCollection().Remove(orgUnit.ID)
}

func GetAll() *mgo.Query {
	return GetCollection().Find(bson.M{})
}

func SearchAll(search string) *mgo.Query {
	return GetCollection().Find(bson.M{"$text": bson.M{"$search": search}})
}