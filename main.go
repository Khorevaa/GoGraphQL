package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/graphql-go/graphql"
	"github.com/mnmtanish/go-graphiql"
	"github.com/NiciiA/GoGraphQL/domain/type/entityType"
	"github.com/NiciiA/GoGraphQL/domain/model/entityModel"
	"gopkg.in/mgo.v2/bson"
	"github.com/NiciiA/GoGraphQL/domain/type/groupType"
	"github.com/NiciiA/GoGraphQL/domain/model/groupModel"
	"github.com/NiciiA/GoGraphQL/domain/type/priorityType"
	"github.com/NiciiA/GoGraphQL/domain/type/categoryType"
	"github.com/NiciiA/GoGraphQL/domain/type/tagType"
	"github.com/NiciiA/GoGraphQL/domain/type/contactType"
	"github.com/NiciiA/GoGraphQL/domain/type/newsType"
	"github.com/NiciiA/GoGraphQL/domain/type/orgUnitType"
	"github.com/NiciiA/GoGraphQL/dataaccess/entityDao"
	"github.com/NiciiA/GoGraphQL/domain/type/activityType"
	"github.com/NiciiA/GoGraphQL/domain/type/fileType"
)

var (
	Schema graphql.Schema
)

/*
	News TODO - REST Client
	Permission / Roles

	And Account Managment with REST

 */
func init() {
	mutationType := graphql.NewObject(graphql.ObjectConfig{
		Name: "Mutation",
		Fields: graphql.Fields{
			"createCategory": &graphql.Field{
				Type: categoryType.Type,
			},
			"updateCategory": &graphql.Field{
				Type: categoryType.Type,
			},
			"removeCategory": &graphql.Field{
				Type: categoryType.Type,
			},
			"disableCategory": &graphql.Field{
				Type: categoryType.Type,
			},
			"createContact": &graphql.Field{
				Type: contactType.Type,
			},
			"updateContact": &graphql.Field{
				Type: contactType.Type,
			},
			"removeContact": &graphql.Field{
				Type: contactType.Type,
			},
			"disableContact": &graphql.Field{
				Type: contactType.Type,
			},
			"createEntity": &graphql.Field{
				Type: entityType.Type,
			},
			"updateEntity": &graphql.Field{
				Type: entityType.Type,
			},
			"removeEntity": &graphql.Field{
				Type: entityType.Type,
			},
			"disableEntity": &graphql.Field{
				Type: entityType.Type,
			},
			"pushEntityFile": &graphql.Field{
				Type: fileType.Type,
			},
			"removeEntityFile": &graphql.Field{
				Type: fileType.Type,
			},
			"pushEntityActivity": &graphql.Field{
				Type: activityType.Type,
			},
			"removeEntityActivity": &graphql.Field{
				Type: activityType.Type,
			},
			"createGroup": &graphql.Field{
				Type: groupType.Type,
			},
			"updateGroup": &graphql.Field{
				Type: groupType.Type,
			},
			"removeGroup": &graphql.Field{
				Type: groupType.Type,
			},
			"disableGroup": &graphql.Field{
				Type: groupType.Type,
			},
			"createNews": &graphql.Field{
				Type: newsType.Type,
			},
			"updateNews": &graphql.Field{
				Type: newsType.Type,
			},
			"removeNews": &graphql.Field{
				Type: newsType.Type,
			},
			"disableNews": &graphql.Field{
				Type: newsType.Type,
			},
			"pushNewsFile": &graphql.Field{
				Type: fileType.Type,
			},
			"removeNewsFile": &graphql.Field{
				Type: fileType.Type,
			},
			"removeNewsComment": &graphql.Field{
				Type: activityType.Type,
			},
			"pushNewsComment": &graphql.Field{
				Type: activityType.Type,
			},
			"createOrgUnit": &graphql.Field{
				Type: orgUnitType.Type,
			},
			"updateOrgUnit": &graphql.Field{
				Type: orgUnitType.Type,
			},
			"removeOrgUnit": &graphql.Field{
				Type: orgUnitType.Type,
			},
			"disableOrgUnit": &graphql.Field{
				Type: orgUnitType.Type,
			},
			"createPriority": &graphql.Field{
				Type: priorityType.Type,
			},
			"updatePriority": &graphql.Field{
				Type: priorityType.Type,
			},
			"removePriority": &graphql.Field{
				Type: priorityType.Type,
			},
			"disablePriority": &graphql.Field{
				Type: priorityType.Type,
			},
			"createTag": &graphql.Field{
				Type: tagType.Type,
			},
			"updateTag": &graphql.Field{
				Type: tagType.Type,
			},
			"removeTag": &graphql.Field{
				Type: tagType.Type,
			},
			"disableTag": &graphql.Field{
				Type: tagType.Type,
			},
		},
	})
	queryType := graphql.NewObject(graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"categoryList": &graphql.Field{
				Type: categoryType.Type,
				Args: graphql.FieldConfigArgument{
					"type": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					// type == news oder entity
					// curl -g 'http://localhost:8080/graphql?query={allTickets{id}}'
					return groupModel.Group{ID: bson.ObjectIdHex("x")}, nil
				},
			},
			"contactList": &graphql.Field{
				Type: contactType.Type,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					// curl -g 'http://localhost:8080/graphql?query={allTickets{id}}'
					return entityModel.Entity{CreatedDate: "fgdfgdfgfdg", Disabled: false, Groups: []string{"customer", "internal"}}, nil
				},
			},
			"entityList": &graphql.Field{
				Type: graphql.NewList(entityType.Type),
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					entityList := []entityModel.Entity{}
					return entityDao.GetAll().All(&entityList), nil
					// curl -g 'http://localhost:8080/graphql?query={allTickets{id}}'
					// return entityModel.Entity{CreatedDate: "fgdfgdfgfdg", Disabled: false, Groups: []string{"customer", "internal"}}, nil
				},
			},
			"groupList": &graphql.Field{
				Type: groupType.Type,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					// curl -g 'http://localhost:8080/graphql?query={allTickets{id}}'
					return groupModel.Group{ID: bson.ObjectIdHex("x")}, nil
				},
			},
			"newsList": &graphql.Field{
				Type: newsType.Type,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					// curl -g 'http://localhost:8080/graphql?query={allTickets{id}}'
					return groupModel.Group{ID: bson.ObjectIdHex("x")}, nil
				},
			},
			"orgUnitList": &graphql.Field{
				Type: orgUnitType.Type,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					// curl -g 'http://localhost:8080/graphql?query={allTickets{id}}'
					return groupModel.Group{ID: bson.ObjectIdHex("x")}, nil
				},
			},
			"priorityList": &graphql.Field{
				Type: priorityType.Type,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					// curl -g 'http://localhost:8080/graphql?query={allTickets{id}}'
					return groupModel.Group{ID: bson.ObjectIdHex("x")}, nil
				},
			},
			"tagList": &graphql.Field{
				Type: tagType.Type,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					// curl -g 'http://localhost:8080/graphql?query={allTickets{id}}'
					return groupModel.Group{ID: bson.ObjectIdHex("x")}, nil
				},
			},
			"category": &graphql.Field{
				Type: categoryType.Type,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					idQuery, _ := p.Args["id"].(string)
					if !bson.IsObjectIdHex(idQuery) {

					}
					// curl -g 'http://localhost:8080/graphql?query={allTickets{id}}'
					return entityModel.Entity{ID: bson.ObjectIdHex(idQuery), CreatedDate: "fgdfgdfgfdg", Disabled: false, Groups: []string{"customer", "internal"}}, nil
				},
			},
			"contact": &graphql.Field{
				Type: contactType.Type,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					idQuery, _ := p.Args["id"].(string)
					if !bson.IsObjectIdHex(idQuery) {

					}
					// curl -g 'http://localhost:8080/graphql?query={allTickets{id}}'
					return entityModel.Entity{ID: bson.ObjectIdHex(idQuery), CreatedDate: "fgdfgdfgfdg", Disabled: false, Groups: []string{"customer", "internal"}}, nil
				},
			},
			"entity": &graphql.Field{
				Type: entityType.Type,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					// curl -g 'http://localhost:8080/graphql?query={allTickets{id}}'
					// return entityModel.Entity{ID: bson.ObjectIdHex(idQuery), CreatedDate: "fgdfgdfgfdg", Disabled: false, Groups: []string{"customer", "internal"}}, nil
					idQuery, _ := p.Args["id"].(string)
					if !bson.IsObjectIdHex(idQuery) {
						return nil, nil
					}
					entity := entityModel.Entity{}
					return entityDao.GetById(bson.ObjectIdHex(idQuery)).One(&entity), nil
				},
			},
			"group": &graphql.Field{
				Type: groupType.Type,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					idQuery, _ := p.Args["id"].(string)
					if !bson.IsObjectIdHex(idQuery) {

					}
					// curl -g 'http://localhost:8080/graphql?query={allTickets{id}}'
					return entityModel.Entity{ID: bson.ObjectIdHex(idQuery), CreatedDate: "fgdfgdfgfdg", Disabled: false, Groups: []string{"customer", "internal"}}, nil
				},
			},
			"news": &graphql.Field{
				Type: newsType.Type,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					idQuery, _ := p.Args["id"].(string)
					if !bson.IsObjectIdHex(idQuery) {

					}
					// curl -g 'http://localhost:8080/graphql?query={allTickets{id}}'
					return entityModel.Entity{ID: bson.ObjectIdHex(idQuery), CreatedDate: "fgdfgdfgfdg", Disabled: false, Groups: []string{"customer", "internal"}}, nil
				},
			},
			"orgUnit": &graphql.Field{
				Type: orgUnitType.Type,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					idQuery, _ := p.Args["id"].(string)
					if !bson.IsObjectIdHex(idQuery) {

					}
					// curl -g 'http://localhost:8080/graphql?query={allTickets{id}}'
					return entityModel.Entity{ID: bson.ObjectIdHex(idQuery), CreatedDate: "fgdfgdfgfdg", Disabled: false, Groups: []string{"customer", "internal"}}, nil
				},
			},
			"priority": &graphql.Field{
				Type: priorityType.Type,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					idQuery, _ := p.Args["id"].(string)
					if !bson.IsObjectIdHex(idQuery) {

					}
					// curl -g 'http://localhost:8080/graphql?query={allTickets{id}}'
					return entityModel.Entity{ID: bson.ObjectIdHex(idQuery), CreatedDate: "fgdfgdfgfdg", Disabled: false, Groups: []string{"customer", "internal"}}, nil
				},
			},
			"tag": &graphql.Field{
				Type: tagType.Type,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					idQuery, _ := p.Args["id"].(string)
					if !bson.IsObjectIdHex(idQuery) {

					}
					// curl -g 'http://localhost:8080/graphql?query={allTickets{id}}'
					return entityModel.Entity{ID: bson.ObjectIdHex(idQuery), CreatedDate: "fgdfgdfgfdg", Disabled: false, Groups: []string{"customer", "internal"}}, nil
				},
			},
		},
	})
	Schema, _ = graphql.NewSchema(graphql.SchemaConfig{
		Query: queryType,
		Mutation: mutationType,
	})
}

func main() {
	http.HandleFunc("/graphql", serveGraphQL(Schema))
	http.HandleFunc("/", graphiql.ServeGraphiQL)
	fmt.Println("Now server is running on port 8080")
	http.ListenAndServe(":8080", nil)
}

func serveGraphQL(s graphql.Schema) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		sendError := func(err error) {
			w.WriteHeader(500)
			w.Write([]byte(err.Error()))
		}

		req := &graphiql.Request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			sendError(err)
			return
		}

		res := graphql.Do(graphql.Params{
			Schema:        s,
			RequestString: req.Query,
		})

		if err := json.NewEncoder(w).Encode(res); err != nil {
			sendError(err)
		}
	}
}
