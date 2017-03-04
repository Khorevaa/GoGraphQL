package contactService

import (
	"github.com/NiciiA/GoGraphQL/domain/model"
	"github.com/NiciiA/GoGraphQL/webapp/authHandler"
	"github.com/NiciiA/GoGraphQL/dataaccess/contactDao"
	"gopkg.in/mgo.v2/bson"
)

func Create(contact model.Contact) {
	if authHandler.CurrentAccount.Roles[0] == "customer" || authHandler.CurrentAccount.Roles[0] == "administrator" {
		contactDao.Insert(contact)
	}
}

func Update(preContact model.Contact, contact model.Contact) {
	if authHandler.CurrentAccount.Roles[0] == "administrator"{
		contactDao.Update(contact)
	}
	if authHandler.CurrentAccount.Roles[0] == "customer"{
		if bson.ObjectIdHex(preContact.Accounts[0]) == authHandler.CurrentAccount.ID {
			contactDao.Update(contact)
		}
	}
}

func Remove(contact model.Contact) {
	if authHandler.CurrentAccount.Roles[0] == "administrator"{
		contactDao.Delete(contact)
	}
	if authHandler.CurrentAccount.Roles[0] == "customer"{
		if bson.ObjectIdHex(contact.Accounts[0]) == authHandler.CurrentAccount.ID {
			contactDao.Delete(contact)
		}
	}
}