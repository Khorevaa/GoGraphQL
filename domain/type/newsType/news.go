package newsType

import (
	"github.com/graphql-go/graphql"
	"github.com/NiciiA/GoGraphQL/domain/type/tagType"
	"github.com/NiciiA/GoGraphQL/domain/type/groupType"
	"github.com/NiciiA/GoGraphQL/domain/type/categoryType"
	"github.com/NiciiA/GoGraphQL/domain/model/newsModel"
	"github.com/NiciiA/GoGraphQL/domain/type/contactType"
	"github.com/NiciiA/GoGraphQL/domain/type/fileType"
)

var Type *graphql.Object = graphql.NewObject(graphql.ObjectConfig{
	Name:        "News",
	Description: "A News response",
	Fields: graphql.Fields{
		"_id": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if news, ok := p.Source.(newsModel.News); ok {
					return news.ID.Hex(), nil
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
		"title": &graphql.Field{
			Type: graphql.String,
		},
		"text": &graphql.Field{
			Type: graphql.String,
		},
		"intern": &graphql.Field{
			Type: graphql.Boolean,
		},
		"tags": &graphql.Field{
			Type: graphql.NewList(tagType.Type),
			Description: "The tags of the news.",
		},
		"groups": &graphql.Field{
			Type: graphql.NewList(groupType.Type),
			Description: "The groups of the news.",
		},
		"important": &graphql.Field{
			Type: graphql.Boolean,
			Description: "Is the news important ?",
		},
		"category": &graphql.Field{
			Type: categoryType.Type,
			Description: "The category of the news.",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if news, ok := p.Source.(newsModel.News); ok {
					return news.Category, nil
				}
				return nil, nil
			},
		},
		"creator": &graphql.Field{
			Type: contactType.Type,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if news, ok := p.Source.(newsModel.News); ok {
					return news.Creator, nil
				}
				return nil, nil
			},
		},
		"files": &graphql.Field{
			Type: graphql.NewList(fileType.Type),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if news, ok := p.Source.(newsModel.News); ok {
					return []string{news.ID.Hex()}, nil
				}
				return nil, nil
			},
		},
	},
})