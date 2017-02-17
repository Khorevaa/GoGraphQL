package orgUnitDao

import (
"github.com/NiciiA/GoGraphQL/dataaccess/mongoAccess"
"github.com/NiciiA/GoGraphQL/domain/model/orgUnitModel"
"gopkg.in/mgo.v2"
"gopkg.in/mgo.v2/bson"
)

var collectionName = "orgUnit"

var session *mgo.Session

var OrgUnitList map[string]orgUnitModel.OrgUnit = make(map[string]orgUnitModel.OrgUnit)

func GetCollection() *mgo.Collection {
	return mongoAccess.GetCollection(session, collectionName)
}

func GetByKey(key string) orgUnitModel.OrgUnit {
	return OrgUnitList[key]
}

func AddOrgUnit(c orgUnitModel.OrgUnit) {
	OrgUnitList[c.Name] = c
	Insert(c)
}

func UpdateOrgUnit(c orgUnitModel.OrgUnit) {
	OrgUnitList[c.Name] = c
	Update(c)
}

func init() {
	session = mongoAccess.Session.Clone()
	var orgUnitList []orgUnitModel.OrgUnit
	GetAll().All(&orgUnitList)
	for _, orgU := range orgUnitList {
		AddOrgUnit(orgU)
	}
}

func Insert(orgUnit orgUnitModel.OrgUnit) {
	GetCollection().Insert(&orgUnit)
}

func Update(orgUnit orgUnitModel.OrgUnit) {
	GetCollection().Update(orgUnit.ID, &orgUnit)
}

func Delete(orgUnit orgUnitModel.OrgUnit) {
	delete(OrgUnitList, orgUnit.Name)
	GetCollection().Remove(orgUnit.ID)
}

func GetAll() *mgo.Query {
	return GetCollection().Find(bson.M{})
}