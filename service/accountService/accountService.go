package accountService

import (
	"github.com/NiciiA/GoGraphQL/domain/model"
	"github.com/NiciiA/GoGraphQL/webapp/authHandler"
	"github.com/NiciiA/GoGraphQL/dataaccess/accountDao"
)

func Create(account model.Account) {
	if authHandler.CurrentAccount.Roles[0] == "customer" || authHandler.CurrentAccount.Roles[0] == "administrator" {
		accountDao.Insert(account)
	}
}

func Update(preAccount model.Account, account model.Account) {
	if authHandler.CurrentAccount.Roles[0] == "administrator"{
		accountDao.Update(account)
	}
	if authHandler.CurrentAccount.Roles[0] == "customer"{
		if preAccount.ID == authHandler.CurrentAccount.ID {
			accountDao.Update(account)
		}
	}
}

func Remove(account model.Account) {
	if authHandler.CurrentAccount.Roles[0] == "administrator"{
		accountDao.Delete(account)
	}
	if authHandler.CurrentAccount.Roles[0] == "customer"{
		if account.ID == authHandler.CurrentAccount.ID {
			accountDao.Delete(account)
		}
	}
}