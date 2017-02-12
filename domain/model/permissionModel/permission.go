package permissionModel

import "gopkg.in/mgo.v2/bson"

type Permissions struct {
	ID bson.ObjectId `bson:"_id" json:"_id"`
	Disabled bool `bson:"disabled" json:"disabled"`
	CreatedDate string `bson:"createdDate" json:"createdDate"`
	ModifiedDate string `bson:"modifiedDate" json:"modifiedDate"`
	Name string `bson:"name" json:"name"`
	TicketPermission TicketPermission `bson:"ticketPermission" json:"ticketPermission"`
	NewsPermission NewsPermission `bson:"newsPermission" json:"newsPermission"`
	UsertaskPermission UsertaskPermission `bson:"usertaskPermission" json:"usertaskPermission"`
	GeneralPermission GeneralPermission `bson:"generalPermission" json:"generalPermission"`
}

type TicketPermission struct {
	Create bool `bson:"create" json:"create"`
	Read bool `bson:"read" json:"read"`
	Edit bool `bson:"edit" json:"edit"`
}

type NewsPermission struct {
	Create bool `bson:"create" json:"create"`
	Read bool `bson:"read" json:"read"`
	Edit bool `bson:"edit" json:"edit"`
}

type UsertaskPermission struct {
	Edit bool `bson:"edit" json:"edit"`
}

type GeneralPermission struct {
	Admin bool `bson:"admin" json:"admin"`
	Internal bool `bson:"internal" json:"internal"`
}