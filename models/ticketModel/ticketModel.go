package ticketModel

import (
	"github.com/graphql-go/graphql"
)


var (
	ticketType *graphql.Object
)

type Ticket  struct {
	ID int
}

func init() {

}