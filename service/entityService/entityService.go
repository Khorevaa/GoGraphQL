package entityService

import (
	"github.com/NiciiA/GoGraphQL/domain/model"
	"github.com/NiciiA/GoGraphQL/webapp/authHandler"
	"github.com/NiciiA/GoGraphQL/dataaccess/entityDao"
)

func Create(entity model.Entity) {
	if authHandler.CurrentAccount.Roles[0] == "customer" || authHandler.CurrentAccount.Roles[0] == "administrator" {
		entityDao.Insert(entity)
	}
}