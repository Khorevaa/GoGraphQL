package entityService

import (
	"github.com/NiciiA/GoGraphQL/domain/model"
	"github.com/NiciiA/GoGraphQL/webapp/authHandler"
	"github.com/NiciiA/GoGraphQL/dataaccess/entityDao"
	"gopkg.in/mgo.v2/bson"
)

func Create(entity model.Entity) {
	if authHandler.CurrentAccount.Roles[0] == "customer" || authHandler.CurrentAccount.Roles[0] == "administrator" {
		entityDao.Insert(entity)
	}
}

func Update(preEntity model.Entity, entity model.Entity) {
	if authHandler.CurrentAccount.Roles[0] == "administrator"{
		entityDao.Update(entity)
	}
	if authHandler.CurrentAccount.Roles[0] == "customer"{
		if bson.ObjectIdHex(preEntity.Creator) == authHandler.CurrentAccount.ID {
			entityDao.Update(entity)
		}
	}
}

func Remove(entity model.Entity) {
	if authHandler.CurrentAccount.Roles[0] == "administrator"{
		entityDao.Delete(entity)
	}
	if authHandler.CurrentAccount.Roles[0] == "customer"{
		if bson.ObjectIdHex(entity.Creator) == authHandler.CurrentAccount.ID {
			entityDao.Delete(entity)
		}
	}
}