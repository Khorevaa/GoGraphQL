package fileType

import (
	"github.com/graphql-go/graphql"
	"github.com/NiciiA/GoGraphQL/domain/model/fileModel"
)

var Type *graphql.Object = graphql.NewObject(graphql.ObjectConfig{
	Name:        "File",
	Description: "A File response",
	Fields: graphql.Fields{
		"_id": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if file, ok := p.Source.(fileModel.File); ok {
					return file.ID.Hex(), nil
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
		"creator": &graphql.Field{
			Type: graphql.String,
		},
		"referenceClass": &graphql.Field{
			Type: graphql.String,
		},
		"referenceId": &graphql.Field{
			Type: graphql.String,
		},
		"filePath": &graphql.Field{
			Type: graphql.String,
		},
		"fileName": &graphql.Field{
			Type: graphql.String,
		},
		"fileContent": &graphql.Field{
			Type: graphql.String,
		},
	},
})