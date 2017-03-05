package newsActivityService

import (
	"github.com/NiciiA/GoGraphQL/domain/model"
	"github.com/NiciiA/GoGraphQL/webapp/authHandler"
	"gopkg.in/mgo.v2/bson"
	"github.com/NiciiA/GoGraphQL/dataaccess/newsActivityDao"
)

func Create(activity model.Activity) {
	if authHandler.CurrentAccount.Roles[0] == "administrator" || authHandler.CurrentAccount.Roles[0] == "customer" {
		newsActivityDao.Insert(activity)
	}
}

func Remove(news model.News, activity model.Activity) {
	if authHandler.CurrentAccount.Roles[0] == "administrator"{
		newsActivityDao.Delete(activity)
	}
	if authHandler.CurrentAccount.Roles[0] == "customer"{
		if bson.ObjectIdHex(news.Creator) == authHandler.CurrentAccount.ID {
			newsActivityDao.Delete(activity)
		}
	}
}