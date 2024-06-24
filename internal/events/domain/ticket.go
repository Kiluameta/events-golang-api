package domain

import "errors"

type TicketType string

var (
	ErrTicketPriceGreater = errors.New("ticket price must be greater than zero")
)

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

func IsValidTicketType(ticketType TicketType) bool {
	return ticketType == TicketTypeFull || ticketType == TicketTypeHalf
}

func (t *Ticket) CalculatePrice() {
	if t.TicketType == TicketTypeHalf {
		t.Price /= 2
	}
}

func (t *Ticket) Validate() error {
	if t.Price <= 0 {
		return ErrTicketPriceGreater
	}
	return nil
}