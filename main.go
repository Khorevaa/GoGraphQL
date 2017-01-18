package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/graphql-go/graphql"
	"github.com/NiciiA/GoGraphQL/services/ticketService"
	"github.com/NiciiA/GoGraphQL/types/ticketType"
)

var (
	Schema graphql.Schema
	ticketServiceVar = ticketService.TicketService{}
	ticketTypeVar = ticketType.TicketType{}
)

func init() {
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
		},
	})
	Schema, _ = graphql.NewSchema(graphql.SchemaConfig{
		Query: queryType,
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
