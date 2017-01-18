package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/graphql-go/graphql"
	"github.com/NiciiA/GoGraphQL/models/ticketModel"
	"github.com/NiciiA/GoGraphQL/services/ticketService"
)

var (
	Schema graphql.Schema
	ticketType *graphql.Object
)

var ticketServiceVar = ticketService.TicketService{}

func init() {
	ticketType = graphql.NewObject(graphql.ObjectConfig{
		Name:        "Ticket",
		Description: "A Ticket",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.String),
				Description: "The id of the ticket.",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if ticket, ok := p.Source.(ticketModel.Ticket); ok {
						return ticket.ID, nil
					}
					return nil, nil
				},
			},
		},
	})
	queryType := graphql.NewObject(graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"allTickets": &graphql.Field{
				Type: graphql.NewList(ticketType),
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
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
