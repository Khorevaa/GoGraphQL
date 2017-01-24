package entityTypesService

import "github.com/NiciiA/GoGraphQL/models/entityTypeModel"

type EntityTypesService struct {

}

func SelectAllStatement() string {
	return "SELECT * FROM EntityTypes"
}

func SelectAllUndisabledStatement() string {
	return SelectAllStatement() + " " + "WHERE disabled = 0"
}

func (e EntityTypesService) GetAll() []entityTypesModel.EntityType {
	var entityTypes []entityTypesModel.EntityType
	return entityTypes
}