package model

import "gopkg.in/mgo.v2/bson"

type Account  struct {
	ID bson.ObjectId `bson:"_id,omitempty" json:"_id"`
	Disabled bool `bson:"disabled" json:"disabled"`
	CreatedDate string `bson:"createdDate" json:"createdDate"`
	ModifiedDate string `bson:"modifiedDate" json:"modifiedDate"`
	UserName string `bson:"userName" json:"userName"`
	Password string `bson:"password" json:"password"`
	EMail string `bson:"eMail" json:"eMail"`
	Phone string `bson:"phone" json:"phone"`
	Roles []string `bson:"roles" json:"roles"`
	Groups []string `bson:"groups" json:"groups"`
}