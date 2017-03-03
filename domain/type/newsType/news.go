package newsType

import (
	"github.com/graphql-go/graphql"
	"github.com/NiciiA/GoGraphQL/domain/type/tagType"
	"github.com/NiciiA/GoGraphQL/domain/type/groupType"
	"github.com/NiciiA/GoGraphQL/domain/type/categoryType"
	"github.com/NiciiA/GoGraphQL/domain/type/contactType"
	"github.com/NiciiA/GoGraphQL/domain/type/fileType"
	"github.com/NiciiA/GoGraphQL/dataaccess/groupDao"
	"gopkg.in/mgo.v2/bson"
	"github.com/NiciiA/GoGraphQL/dataaccess/tagDao"
	"github.com/NiciiA/GoGraphQL/dataaccess/categoryDao"
	"github.com/NiciiA/GoGraphQL/dataaccess/contactDao"
	"github.com/NiciiA/GoGraphQL/domain/model"
)

var Type *graphql.Object = graphql.NewObject(graphql.ObjectConfig{
	Name:        "News",
	Description: "A News response",
	Fields: graphql.Fields{
		"_id": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if news, ok := p.Source.(model.News); ok {
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
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if news, ok := p.Source.(model.News); ok {
					tagList := []model.Tag{}
					for _, tagId := range news.Tags {
						tagList = append(tagList, tagDao.GetByKey(bson.ObjectIdHex(tagId)))
					}
					return tagList, nil
				}
				return nil, nil
			},
		},
		"groups": &graphql.Field{
			Type: graphql.NewList(groupType.Type),
			Description: "The groups of the news.",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if news, ok := p.Source.(model.News); ok {
					groupList := []model.Group{}
					for _, groupId := range news.Groups {
						groupList = append(groupList, groupDao.GetByKey(bson.ObjectIdHex(groupId)))
					}
					return groupList, nil
				}
				return nil, nil
			},
		},
		"important": &graphql.Field{
			Type: graphql.Boolean,
			Description: "Is the news important ?",
		},
		"category": &graphql.Field{
			Type: categoryType.Type,
			Description: "The category of the news.",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if news, ok := p.Source.(model.News); ok {
					return categoryDao.GetByKey(bson.ObjectIdHex(news.Category)), nil
				}
				return nil, nil
			},
		},
		"creator": &graphql.Field{
			Type: contactType.Type,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if news, ok := p.Source.(model.News); ok {
					return contactDao.GetById(bson.ObjectIdHex(news.Creator)), nil
				}
				return nil, nil
			},
		},
		"files": &graphql.Field{
			Type: fileType.Type,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if news, ok := p.Source.(model.News); ok {
					return news.ID.Hex(), nil
				}
				return nil, nil
			},
		},
	},
})