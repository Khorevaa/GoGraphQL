package tagType

import "github.com/graphql-go/graphql"

var Type *graphql.Object = graphql.NewObject(graphql.ObjectConfig{
	Name:        "Tag",
	Description: "A Tag response",
	Fields: graphql.Fields{
		"_id": &graphql.Field{
			Type: graphql.NewNonNull(graphql.String),
		},
		"name": &graphql.Field{
			Type: graphql.NewNonNull(graphql.String),
		},
		"style": &graphql.Field{
			Type: graphql.NewNonNull(graphql.String),
		},
	},
})