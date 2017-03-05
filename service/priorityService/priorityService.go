package priorityService

import (
	"github.com/NiciiA/GoGraphQL/domain/model"
	"github.com/NiciiA/GoGraphQL/webapp/authHandler"
	"github.com/NiciiA/GoGraphQL/dataaccess/priorityDao"
)

func Create(priority model.Priority) {
	if authHandler.CurrentAccount.Roles[0] == "administrator" {
		priorityDao.AddPriority(priority)
	}
}

func Update(priority model.Priority) {
	if authHandler.CurrentAccount.Roles[0] == "administrator"{
		priorityDao.UpdatePriority(priority)
	}
}

func Remove(priority model.Priority) {
	if authHandler.CurrentAccount.Roles[0] == "administrator"{
		priorityDao.Delete(priority)
	}
}