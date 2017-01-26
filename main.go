package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/graphql-go/graphql"
	"github.com/NiciiA/GoGraphQL/services/ticketService"
	"github.com/NiciiA/GoGraphQL/types/ticketType"
	"github.com/NiciiA/GoGraphQL/types/entityTypesType"
	"github.com/NiciiA/GoGraphQL/services/entityTypesService"
	"github.com/NiciiA/GoGraphQL/models/entityTypeModel"
	"time"
	"github.com/mnmtanish/go-graphiql"
)

var (
	Schema graphql.Schema
)

func init() {
	mutationType := graphql.NewObject(graphql.ObjectConfig{
		Name: "Mutation",
		Fields: graphql.Fields{
			"createEntityType": &graphql.Field{
				Type: entityTypesType.Type,
				Args: graphql.FieldConfigArgument{
					"key": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"name": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"type": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"disabled": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Boolean),
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {

					key, _ := params.Args["key"].(string)
					name, _ := params.Args["name"].(string)
					entityType, _ := params.Args["type"].(string)
					disabled, _ := params.Args["disabled"].(bool)

					var newEntity entityTypesModel.EntityType = entityTypesModel.EntityType{ID: "", Key: key, Name: name, Type: entityType, Disabled: disabled, ClassName: "com.projectaton.entityTypes." + key, CreatedDate: time.Now().String(), ModifiedDate: time.Now().String()}
					newEntity.Create()

					return newEntity, nil
				},
			},
		},
	})
	queryType := graphql.NewObject(graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"allTickets": &graphql.Field{
				Type: graphql.NewList(ticketType.Type),
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					// curl -g 'http://localhost:8080/graphql?query={allTickets{id}}'
					return ticketService.GetAll(), nil
				},
			},
			"allEntityTypes": &graphql.Field{
				Type: graphql.NewList(entityTypesType.Type),
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					// curl -g 'http://localhost:8080/graphql?query={allEntityTypes{id,key,name,className,type,disabled,createdDate,modifiedDate}}'
					return entityTypesService.GetAll(), nil
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
