package entityType

import (
	"github.com/graphql-go/graphql"
	"github.com/NiciiA/GoGraphQL/domain/type/groupType"
	"github.com/NiciiA/GoGraphQL/domain/type/tagType"
	"github.com/NiciiA/GoGraphQL/domain/type/priorityType"
	"github.com/NiciiA/GoGraphQL/domain/type/categoryType"
	"github.com/NiciiA/GoGraphQL/domain/model/entityModel"
	"github.com/NiciiA/GoGraphQL/domain/type/contactType"
	"github.com/NiciiA/GoGraphQL/domain/type/fileType"
	"github.com/NiciiA/GoGraphQL/dataaccess/categoryDao"
	"gopkg.in/mgo.v2/bson"
)

var Type *graphql.Object = graphql.NewObject(graphql.ObjectConfig{
	Name:        "Entity",
	Description: "A Entity response",
	Fields: graphql.Fields{
		"_id": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if entity, ok := p.Source.(entityModel.Entity); ok {
					return entity.ID.Hex(), nil
				}
				return nil, nil
			},
		},
		"disabled": &graphql.Field{
			Type: graphql.Boolean,
		},
		"createdDate": &graphql.Field{
			Type: graphql.String,
		},
		"modifiedDate": &graphql.Field{
			Type: graphql.String,
		},
		"subject": &graphql.Field{
			Type: graphql.String,
		},
		"description": &graphql.Field{
			Type: graphql.String,
		},
		"tags": &graphql.Field{
			Type: graphql.NewList(tagType.Type),
			Description: "The tags of the entity.",
		},
		"groups": &graphql.Field{
			Type: graphql.NewList(groupType.Type),
			Description: "The groups of the entity.",
		},
		"priority": &graphql.Field{
			Type: priorityType.Type,
			Description: "The priority of the entity.",
		},
		"category": &graphql.Field{
			Type: categoryType.Type,
			Description: "The category of the entity.",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if entity, ok := p.Source.(entityModel.Entity); ok {
					return categoryDao.GetByKey(bson.ObjectIdHex(entity.Category)), nil
				}
				return nil, nil
			},
		},
		"creator": &graphql.Field{
			Type: contactType.Type,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if entity, ok := p.Source.(entityModel.Entity); ok {
					return entity.Creator, nil
				}
				return nil, nil
			},
		},
		"files": &graphql.Field{
			Type: fileType.Type,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if entity, ok := p.Source.(entityModel.Entity); ok {
					return entity.ID.Hex(), nil
				}
				return nil, nil
			},
		},
	},
})