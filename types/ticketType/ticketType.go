package ticketType

import (
	"github.com/graphql-go/graphql"
	"github.com/NiciiA/GoGraphQL/models/ticketModel"
)

type TicketType struct { }

func (t TicketType) GetType() *graphql.Object  {
	return graphql.NewObject(graphql.ObjectConfig{
		Name:        "Ticket",
		Description: "A Ticket response",
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
}