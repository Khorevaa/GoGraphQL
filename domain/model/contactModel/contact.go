package contactModel

import "gopkg.in/mgo.v2/bson"

type Contact  struct {
	ID bson.ObjectId `bson:"_id" json:"_id"`
	Disabled bool `bson:"disabled" json:"disabled"`
	CreatedDate string `bson:"createdDate" json:"createdDate"`
	ModifiedDate string `bson:"modifiedDate" json:"modifiedDate"`
	FirstName string `bson:"firstName" json:"firstName"`
	LastName string `bson:"lastName" json:"lastName"`
	Street string `bson:"street" json:"street"`
	Village string `bson:"village" json:"village"`
	OrgUnit string `bson:"orgUnit" json:"orgUnit"`
	Accounts []string `bson:"accounts" json:"accounts"`
}