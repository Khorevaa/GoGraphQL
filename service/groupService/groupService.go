package groupService

import (
	"github.com/NiciiA/GoGraphQL/domain/model"
	"github.com/NiciiA/GoGraphQL/webapp/authHandler"
	"github.com/NiciiA/GoGraphQL/dataaccess/groupDao"
)

func Create(group model.Group) {
	if authHandler.CurrentAccount.Roles[0] == "administrator" {
		groupDao.AddGroup(group)
	}
}

func Update(group model.Group) {
	if authHandler.CurrentAccount.Roles[0] == "administrator"{
		groupDao.UpdateGroup(group)
	}
}

func Remove(group model.Group) {
	if authHandler.CurrentAccount.Roles[0] == "administrator"{
		groupDao.Delete(group)
	}
}