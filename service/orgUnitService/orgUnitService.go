package orgUnitService

import (
	"github.com/NiciiA/GoGraphQL/domain/model"
	"github.com/NiciiA/GoGraphQL/webapp/authHandler"
	"github.com/NiciiA/GoGraphQL/dataaccess/orgUnitDao"
)

func Create(orgUnit model.OrgUnit) {
	if authHandler.CurrentAccount.Roles[0] == "administrator" {
		orgUnitDao.AddOrgUnit(orgUnit)
	}
}

func Update(orgUnit model.OrgUnit) {
	if authHandler.CurrentAccount.Roles[0] == "administrator"{
		orgUnitDao.UpdateOrgUnit(orgUnit)
	}
}

func Remove(orgUnit model.OrgUnit) {
	if authHandler.CurrentAccount.Roles[0] == "administrator"{
		orgUnitDao.Delete(orgUnit)
	}
}