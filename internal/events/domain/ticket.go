package domain

type TicketType string

const (
	TicketTypeFull TicketType = "full"
	TicketTypeHalf TicketType = "half"
)

type Ticket struct {
	ID         string
	EventID    string
	Spot       *Spot
	TicketType TicketType
	Price      float64
}