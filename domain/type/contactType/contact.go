package contactType

import (
	"github.com/graphql-go/graphql"
	"github.com/NiciiA/GoGraphQL/domain/model/contactModel"
	"github.com/NiciiA/GoGraphQL/domain/type/orgUnitType"
	"github.com/NiciiA/GoGraphQL/domain/type/fileType"
	"gopkg.in/mgo.v2/bson"
	"github.com/NiciiA/GoGraphQL/dataaccess/orgUnitDao"
)

var Type *graphql.Object = graphql.NewObject(graphql.ObjectConfig{
	Name:        "Contact",
	Description: "A Contact response",
	Fields: graphql.Fields{
		"_id": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if contact, ok := p.Source.(contactModel.Contact); ok {
					return contact.ID.Hex(), nil
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
		"firstName": &graphql.Field{
			Type: graphql.String,
		},
		"lastName": &graphql.Field{
			Type: graphql.String,
		},
		"street": &graphql.Field{
			Type: graphql.String,
		},
		"village": &graphql.Field{
			Type: graphql.String,
		},
		"orgUnit": &graphql.Field{
			Type: orgUnitType.Type,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if contact, ok := p.Source.(contactModel.Contact); ok {
					return orgUnitDao.GetByKey(bson.ObjectIdHex(contact.OrgUnit)), nil
				}
				return nil, nil
			},
		},
		"accounts": &graphql.Field{
			Type: graphql.NewList(graphql.String),
		},
		"profileImage": &graphql.Field{
			Type: fileType.Type,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if contact, ok := p.Source.(contactModel.Contact); ok {
					return contact.ID.Hex(), nil
				}
				return nil, nil
			},
		},
	},
})