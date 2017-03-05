package tagService

import (
	"github.com/NiciiA/GoGraphQL/domain/model"
	"github.com/NiciiA/GoGraphQL/webapp/authHandler"
	"github.com/NiciiA/GoGraphQL/dataaccess/tagDao"
)

func Create(tag model.Tag) {
	if authHandler.CurrentAccount.Roles[0] == "administrator" {
		tagDao.AddTag(tag)
	}
}

func Update(tag model.Tag) {
	if authHandler.CurrentAccount.Roles[0] == "administrator"{
		tagDao.UpdateTag(tag)
	}
}

func Remove(tag model.Tag) {
	if authHandler.CurrentAccount.Roles[0] == "administrator"{
		tagDao.Delete(tag)
	}
}