package groupType

import "github.com/graphql-go/graphql"

var Type *graphql.Object = graphql.NewObject(graphql.ObjectConfig{
	Name:        "Group",
	Description: "A Group response",
	Fields: graphql.Fields{
		"_id": &graphql.Field{
			Type: graphql.NewNonNull(graphql.String),
		},
		"name": &graphql.Field{
			Type: graphql.NewNonNull(graphql.String),
		},
	},
})