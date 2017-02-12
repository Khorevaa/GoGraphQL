package jwtModel

type News  struct {
	JWT string `bson:"jwt" json:"jwt"`
	Account string `bson:"account" json:"account"`
}