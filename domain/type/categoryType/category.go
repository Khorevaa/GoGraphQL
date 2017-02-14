package categoryType

import (
	"github.com/graphql-go/graphql"
	"github.com/NiciiA/GoGraphQL/domain/model/categoryModel"
)

var Type *graphql.Object = graphql.NewObject(graphql.ObjectConfig{
	Name:        "Category",
	Description: "A Category response",
	Fields: graphql.Fields{
		"_id": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if category, ok := p.Source.(categoryModel.Category); ok {
					return category.ID.Hex(), nil
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
		"name": &graphql.Field{
			Type: graphql.String,
		},
		"type": &graphql.Field{
			Type: graphql.String,
		},
	},
})