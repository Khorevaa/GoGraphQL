package accountType

import (
	"github.com/graphql-go/graphql"
	"github.com/NiciiA/GoGraphQL/dataaccess/groupDao"
	"gopkg.in/mgo.v2/bson"
	"github.com/NiciiA/GoGraphQL/domain/type/groupType"
	"github.com/NiciiA/GoGraphQL/domain/model"
)

var Type *graphql.Object = graphql.NewObject(graphql.ObjectConfig{
	Name:        "Account",
	Description: "An Account response",
	Fields: graphql.Fields{
		"_id": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if account, ok := p.Source.(model.Account); ok {
					return account.ID.Hex(), nil
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
		"userName": &graphql.Field{
			Type: graphql.String,
		},
		"eMail": &graphql.Field{
			Type: graphql.String,
		},
		"phone": &graphql.Field{
			Type: graphql.String,
		},
		"roles": &graphql.Field{
			Type: graphql.NewList(graphql.String),
		},
		"groups": &graphql.Field{
			Type: graphql.NewList(groupType.Type),
			Description: "The groups of the news.",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if accountModel, ok := p.Source.(model.Account); ok {
					groupList := []model.Group{}
					for _, groupId := range accountModel.Groups {
						groupList = append(groupList, groupDao.GetByKey(bson.ObjectIdHex(groupId)))
					}
					return groupList, nil
				}
				return nil, nil
			},
		},
	},
})