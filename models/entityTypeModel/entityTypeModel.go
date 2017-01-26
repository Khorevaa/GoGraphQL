package entityTypesModel

import (
	"strconv"
	"database/sql"
	"github.com/NiciiA/GoGraphQL/database/mariaDB"
	"log"
	"fmt"
)

type EntityType  struct {
	ID string
	Key string
	Name string
	ClassName string
	Type string
	Disabled bool
	CreatedDate string
	ModifiedDate string
}

func (e *EntityType) Create() {
	var res sql.Result = mariaDB.Insert(e.Insert())
	id, err := res.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}
	e.ID = strconv.FormatInt(id, 10)
}

func (e EntityType) Insert() string {
	insertString := "INSERT INTO entityTypes" + setString(e) + ", 'createdDate' = CURRENT_TIMESTAMP"
	fmt.Println(insertString)
	return insertString
}

func (e EntityType) Update() string {
	updateString := "UPDATE entityTypes WHERE 'id' = " + e.ID + setString(e)
	return updateString
}

func setString(e EntityType) string {
	setString := " SET 'key' = '" + e.Key + "', 'name' = '" + e.Name + "', 'className' = '" + e.ClassName + "', 'type' = '" + e.Type + "', 'disabled' = " + strconv.FormatBool(e.Disabled) + ", 'modifiedDate' = CURRENT_TIMESTAMP"
	return setString
}