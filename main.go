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
)

var (
	Schema graphql.Schema
	ticketServiceVar = ticketService.TicketService{}
	ticketTypeVar = ticketType.TicketType{}
	entityTypesServiceVar = entityTypesService.EntityTypesService{}
	entityTypesTypeVar = entityTypesType.EntityTypesType{}
)

func init() {
	mutationType := graphql.NewObject(graphql.ObjectConfig{
		Name: "Mutation",
		Fields: graphql.Fields{
			/*
			"entityType": &graphql.Field{
				Type: entityTypesTypeVar.GetType(), // the return type for this field
				Args: graphql.FieldConfigArgument{
					"text": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {

					// marshall and cast the argument value
					text, _ := params.Args["text"].(string)

					// figure out new id
					newID := RandStringRunes(8)

					// perform mutation operation here
					// for e.g. create a Todo and save to DB.
					newTodo := Todo{
						ID:   newID,
						Text: text,
						Done: false,
						AccountID: 1,
					}

					TodoList = append(TodoList, newTodo)

					// return the new Todo object that we supposedly save to DB
					// Note here that
					// - we are returning a `Todo` struct instance here
					// - we previously specified the return Type to be `todoType`
					// - `Todo` struct maps to `todoType`, as defined in `todoType` ObjectConfig`
					return newTodo, nil
				},
			},
			*/
		},
	})
	queryType := graphql.NewObject(graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"allTickets": &graphql.Field{
				Type: graphql.NewList(ticketTypeVar.GetType()),
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					// curl -g 'http://localhost:8080/graphql?query={allTickets{id}}'
					return ticketServiceVar.GetAll(), nil
				},
			},
			"allEntityTypes": &graphql.Field{
				Type: graphql.NewList(entityTypesTypeVar.GetType()),
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					// curl -g 'http://localhost:8080/graphql?query={allTickets{id}}'
					return entityTypesServiceVar.GetAll(), nil
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
	http.HandleFunc("/graphql", func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query()["query"][0]
		result := graphql.Do(graphql.Params{
			Schema: Schema,
			RequestString: query,
		})
		json.NewEncoder(w).Encode(result)
	})
	fmt.Println("Now server is running on port 8080")
	http.ListenAndServe(":8080", nil)
}
