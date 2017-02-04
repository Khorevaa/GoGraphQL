package entityTypesService

import (
	"github.com/NiciiA/GoGraphQL/models/entityTypeModel"
	"github.com/NiciiA/GoGraphQL/database/mariaDB"
	"database/sql"
	"log"
	"fmt"
	"strconv"
)

func SelectAllStatement() string {
	return "SELECT * FROM entityTypes"
}

func SelectAllUndisabledStatement() string {
	return SelectAllStatement() + " " + "WHERE disabled = 0"
}

func GetAll() []entityTypeModel.EntityType {
	var entityTypes []entityTypeModel.EntityType
	var rows *sql.Rows = mariaDB.Select(SelectAllUndisabledStatement())
	for rows.Next() {
		var id int
		var key string
		var name string
		var className string
		var entityType string
		var disabled bool
		var createdDate string
		var modifiedDate string
		if err := rows.Scan(&id, &key, &name, &className, &entityType, &disabled, &createdDate, &modifiedDate); err != nil {
			log.Fatal(err)
			fmt.Println(err)
		}
		entityTypes = append(entityTypes, entityTypeModel.EntityType{ID: strconv.Itoa(id), Key: key, Name: name, ClassName: className, Type: entityType, Disabled: disabled, CreatedDate: createdDate, ModifiedDate: modifiedDate})
	}
	rows.Close()
	return entityTypes
}