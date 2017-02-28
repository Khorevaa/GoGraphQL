package jwtModel

import "github.com/NiciiA/GoGraphQL/domain/model/accountModel"

type JWT struct {
	JWT string `bson:"jwt" json:"jwt"`
	Account accountModel.Account `bson:"account" json:"account"`
}