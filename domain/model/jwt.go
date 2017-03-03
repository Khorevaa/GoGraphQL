package model

type JWT struct {
	JWT string `bson:"jwt" json:"jwt"`
	Account Account `bson:"account" json:"account"`
}