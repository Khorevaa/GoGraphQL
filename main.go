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
)

var (
	Schema graphql.Schema
)

/*
	News TODO - REST Client

	And Account Managment with REST
	Maybe Contact and Org Unit with graphql
 */
func init() {
	mutationType := graphql.NewObject(graphql.ObjectConfig{
		Name: "Mutation",
		Fields: graphql.Fields{
			"createEntity": &graphql.Field{
				Type: entityType.Type,
			},
			"updateEntity": &graphql.Field{
				Type: entityType.Type,
			},
			"removeEntity": &graphql.Field{
				Type: entityType.Type,
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
			"createCategory": &graphql.Field{
				Type: categoryType.Type,
			},
			"updateCategory": &graphql.Field{
				Type: categoryType.Type,
			},
			"removeCategory": &graphql.Field{
				Type: categoryType.Type,
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
			"createTag": &graphql.Field{
				Type: tagType.Type,
			},
			"updateTag": &graphql.Field{
				Type: tagType.Type,
			},
			"removeTag": &graphql.Field{
				Type: tagType.Type,
			},
		},
	})
	queryType := graphql.NewObject(graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"categoryList": &graphql.Field{
				Type: categoryType.Type,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					// curl -g 'http://localhost:8080/graphql?query={allTickets{id}}'
					return groupModel.Group{ID: bson.ObjectIdHex("x")}, nil
				},
			},
			"entityList": &graphql.Field{
				Type: entityType.Type,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					// curl -g 'http://localhost:8080/graphql?query={allTickets{id}}'
					return entityModel.Entity{CreatedDate: "fgdfgdfgfdg", Disabled: false, Groups: []string{"customer", "internal"}}, nil
				},
			},
			"groupList": &graphql.Field{
				Type: groupType.Type,
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
			"entity": &graphql.Field{
				Type: entityType.Type,
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
