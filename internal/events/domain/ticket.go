package domain

import (
	"errors"

	"github.com/google/uuid"
)

type TicketType string

var (
	ErrTicketPriceGreater = errors.New("ticket price must be greater than zero")
	ErrInvalidTicketType = errors.New("invalid ticket type")
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

func NewTicket(event *Event, spot *Spot, ticketType TicketType) (*Ticket, error) {
	if !IsValidTicketType(ticketType) {
		return nil, ErrInvalidTicketType
	}
	
	ticket := &Ticket {
		ID: uuid.New().String(),
		EventID: event.ID,
		Spot: spot,
		TicketType: ticketType,
		Price: event.Price,
	}

	ticket.CalculatePrice()

	if err := ticket.Validate(); err != nil {
		return nil, err
	}
	return ticket, nil
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