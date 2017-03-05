package categoryService

import (
	"github.com/NiciiA/GoGraphQL/domain/model"
	"github.com/NiciiA/GoGraphQL/webapp/authHandler"
	"github.com/NiciiA/GoGraphQL/dataaccess/categoryDao"
)

func Create(category model.Category) {
	if authHandler.CurrentAccount.Roles[0] == "administrator" {
		categoryDao.AddCategory(category)
	}
}

func Update(category model.Category) {
	if authHandler.CurrentAccount.Roles[0] == "administrator"{
		categoryDao.UpdateCategory(category)
	}
}

func Remove(category model.Category) {
	if authHandler.CurrentAccount.Roles[0] == "administrator"{
		categoryDao.Delete(category)
	}
}