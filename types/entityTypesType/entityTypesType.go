package entityTypesType

import (
	"github.com/graphql-go/graphql"
	"github.com/NiciiA/GoGraphQL/models/entityTypeModel"
)

var Type *graphql.Object =  graphql.NewObject(graphql.ObjectConfig{
	Name:        "EntityTypes",
	Description: "A EntityType response",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type:        graphql.NewNonNull(graphql.String),
			Description: "The id of the entityType.",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if entityType, ok := p.Source.(entityTypesModel.EntityType); ok {
					return entityType.ID, nil
				}
				return nil, nil
			},
		},
		"key": &graphql.Field{
			Type:        graphql.NewNonNull(graphql.String),
			Description: "The key of the entityType.",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if entityType, ok := p.Source.(entityTypesModel.EntityType); ok {
					return entityType.Key, nil
				}
				return nil, nil
			},
		},
		"name": &graphql.Field{
			Type:        graphql.NewNonNull(graphql.String),
			Description: "The name of the entityType.",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if entityType, ok := p.Source.(entityTypesModel.EntityType); ok {
					return entityType.Name, nil
				}
				return nil, nil
			},
		},
		"className": &graphql.Field{
			Type:        graphql.NewNonNull(graphql.String),
			Description: "The className of the entityType.",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if entityType, ok := p.Source.(entityTypesModel.EntityType); ok {
					return entityType.ClassName, nil
				}
				return nil, nil
			},
		},
		"type": &graphql.Field{
			Type:        graphql.NewNonNull(graphql.String),
			Description: "The type of the entityType.",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if entityType, ok := p.Source.(entityTypesModel.EntityType); ok {
					return entityType.Type, nil
				}
				return nil, nil
			},
		},
		"disabled": &graphql.Field{
			Type:        graphql.NewNonNull(graphql.Boolean),
			Description: "The disabled of the entityType.",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if entityType, ok := p.Source.(entityTypesModel.EntityType); ok {
					return entityType.Disabled, nil
				}
				return nil, nil
			},
		},
		"createdDate": &graphql.Field{
			Type:        graphql.NewNonNull(graphql.String),
			Description: "The createdDate of the entityType.",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if entityType, ok := p.Source.(entityTypesModel.EntityType); ok {
					return entityType.CreatedDate, nil
				}
				return nil, nil
			},
		},
		"modifiedDate": &graphql.Field{
			Type:        graphql.NewNonNull(graphql.String),
			Description: "The modifiedDate of the entityType.",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if entityType, ok := p.Source.(entityTypesModel.EntityType); ok {
					return entityType.ModifiedDate, nil
				}
				return nil, nil
			},
		},
	},
})