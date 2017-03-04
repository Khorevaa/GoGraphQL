package newsService

import (
	"github.com/NiciiA/GoGraphQL/domain/model"
	"github.com/NiciiA/GoGraphQL/webapp/authHandler"
	"github.com/NiciiA/GoGraphQL/dataaccess/newsDao"
)

func Create(news model.News) {
	if authHandler.CurrentAccount.Roles[0] == "administrator" {
		newsDao.Insert(news)
	}
}

func Update(preNews model.News, news model.News) {
	if authHandler.CurrentAccount.Roles[0] == "administrator"{
		newsDao.Update(news)
	}
}

func Remove(news model.News) {
	if authHandler.CurrentAccount.Roles[0] == "administrator"{
		newsDao.Delete(news)
	}
}