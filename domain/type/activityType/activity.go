package activityType

import (
	"github.com/graphql-go/graphql"
	"github.com/NiciiA/GoGraphQL/domain/model/entityModel"
	"github.com/NiciiA/GoGraphQL/domain/type/contactType"
	"github.com/NiciiA/GoGraphQL/domain/model/activityModel"
)

var Type *graphql.Object = graphql.NewObject(graphql.ObjectConfig{
	Name:        "Activity",
	Description: "An Activity response",
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
		"referenceClass": &graphql.Field{
			Type: graphql.String,
		},
		"referenceId": &graphql.Field{
			Type: graphql.String,
		},
		"comment": &graphql.Field{
			Type: graphql.String,
		},
		"intern": &graphql.Field{
			Type: graphql.Boolean,
		},
		"creator": &graphql.Field{
			Type: contactType.Type,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if activity, ok := p.Source.(activityModel.Activity); ok {
					return activity.Creator, nil
				}
				return nil, nil
			},
		},
	},
})