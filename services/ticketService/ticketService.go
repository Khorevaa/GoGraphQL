package ticketService

import "github.com/NiciiA/GoGraphQL/models/ticketModel"

func GetAll() []ticketModel.Ticket {
	var TicketList []ticketModel.Ticket
	TicketList = append(TicketList, ticketModel.Ticket{ID: "1"}, ticketModel.Ticket{ID: "2"})
	return TicketList
}