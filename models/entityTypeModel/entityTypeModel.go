package entityTypesModel

import (
	"strconv"
)

type EntityType  struct {
	ID string
	Key string
	Name string
	ClassName string
	Type int
	Disabled bool
}

func (e EntityType) Insert() string {
	insertString := "INSERT INTO EntityTypes" + e.ID + e.setString()
	return insertString
}

func (e EntityType) Update() string {
	updateString := "UPDATE EntityTypes WHERE 'id' = " + e.ID + e.setString()
	return updateString
}

func (e EntityType) setString() string {
	setString := " SET 'key' = " + e.Key + " SET 'name' = " + e.Name + " SET 'className' = " + e.ClassName + " SET 'type' = " + strconv.Itoa(e.Type) + " SET 'disabled' = " + strconv.FormatBool(e.Disabled)
	return setString
}