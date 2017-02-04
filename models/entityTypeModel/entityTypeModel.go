package entityTypeModel

import (
	"strconv"
	"database/sql"
	"github.com/NiciiA/GoGraphQL/database/mariaDB"
	"log"
	"fmt"
	"github.com/graphql-go/graphql"
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

var Type *graphql.Object =  graphql.NewObject(graphql.ObjectConfig{
	Name:        "EntityType",
	Description: "A EntityType response",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type:        graphql.NewNonNull(graphql.String),
			Description: "The id of the entityType.",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if entityType, ok := p.Source.(EntityType); ok {
					return entityType.ID, nil
				}
				return nil, nil
			},
		},
		"key": &graphql.Field{
			Type:        graphql.NewNonNull(graphql.String),
			Description: "The key of the entityType.",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if entityType, ok := p.Source.(EntityType); ok {
					return entityType.Key, nil
				}
				return nil, nil
			},
		},
		"name": &graphql.Field{
			Type:        graphql.NewNonNull(graphql.String),
			Description: "The name of the entityType.",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if entityType, ok := p.Source.(EntityType); ok {
					return entityType.Name, nil
				}
				return nil, nil
			},
		},
		"className": &graphql.Field{
			Type:        graphql.NewNonNull(graphql.String),
			Description: "The className of the entityType.",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if entityType, ok := p.Source.(EntityType); ok {
					return entityType.ClassName, nil
				}
				return nil, nil
			},
		},
		"type": &graphql.Field{
			Type:        graphql.NewNonNull(graphql.String),
			Description: "The type of the entityType.",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if entityType, ok := p.Source.(EntityType); ok {
					return entityType.Type, nil
				}
				return nil, nil
			},
		},
		"disabled": &graphql.Field{
			Type:        graphql.NewNonNull(graphql.Boolean),
			Description: "The disabled of the entityType.",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if entityType, ok := p.Source.(EntityType); ok {
					return entityType.Disabled, nil
				}
				return nil, nil
			},
		},
		"createdDate": &graphql.Field{
			Type:        graphql.NewNonNull(graphql.String),
			Description: "The createdDate of the entityType.",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if entityType, ok := p.Source.(EntityType); ok {
					return entityType.CreatedDate, nil
				}
				return nil, nil
			},
		},
		"modifiedDate": &graphql.Field{
			Type:        graphql.NewNonNull(graphql.String),
			Description: "The modifiedDate of the entityType.",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if entityType, ok := p.Source.(EntityType); ok {
					return entityType.ModifiedDate, nil
				}
				return nil, nil
			},
		},
	},
})

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