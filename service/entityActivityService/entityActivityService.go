package entityActivityService

import (
	"github.com/NiciiA/GoGraphQL/domain/model"
	"github.com/NiciiA/GoGraphQL/webapp/authHandler"
	"gopkg.in/mgo.v2/bson"
	"github.com/NiciiA/GoGraphQL/dataaccess/entityActivityDao"
)

func Create(entity model.Entity, activity model.Activity) {
	if authHandler.CurrentAccount.Roles[0] == "administrator" {
		entityActivityDao.Insert(activity)
	}
	if authHandler.CurrentAccount.Roles[0] == "customer" {
		if bson.ObjectIdHex(entity.Creator) == authHandler.CurrentAccount.ID {
			entityActivityDao.Insert(activity)
		}
	}
}

func Remove(entity model.Entity, activity model.Activity) {
	if authHandler.CurrentAccount.Roles[0] == "administrator"{
		entityActivityDao.Delete(activity)
	}
	if authHandler.CurrentAccount.Roles[0] == "customer"{
		if bson.ObjectIdHex(entity.Creator) == authHandler.CurrentAccount.ID {
			entityActivityDao.Delete(activity)
		}
	}
}