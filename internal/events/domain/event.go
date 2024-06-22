package domain

import (
	"errors"
	"time"
)

var (
	ErrEventNameRequired    = errors.New("vent name is required")
	ErrEventDateFuture      = errors.New("vent date must be in the future")
	ErrEventCapacityGreater = errors.New("vent capacity must be greater than zero")
	ErrEventPriceGreater    = errors.New("vent capacity must be price than zero")
)

type Rating string

const (
	RatingFree Rating = "L"
	Rating10    Rating = "L10"
	Rating12    Rating = "L12"
	Rating14    Rating = "L14"
	Rating16    Rating = "L16"
	Rating18    Rating = "L18"
)

type Event struct {
	ID           string
	Name         string
	Location     string
	Organization string
	Rating       Rating
	Date         time.Time
	ImageURL     string
	Capacity     int
	Price        float64
	PartnerID    int
	Spots        []Spot
	Tickets      []Ticket
}

func (e Event) Validate() error {
	validacoes := map[bool]error{
        e.Name == "":              ErrEventNameRequired,
        e.Date.Before(time.Now()): ErrEventDateFuture,
        e.Capacity <= 0:           ErrEventCapacityGreater,
        e.Price <= 0:              ErrEventPriceGreater,
    }

    for condicao, err := range validacoes {
        if condicao {
            return err
        }
    }

    return nil
}

func (e *Event) AddSpot(name string) (*Spot, error) {
	spot, err := NewSpot(e, name)
	if err != nil {
		return nil, err
	}
	e.Spots = append(e.Spots, *spot)
	return spot, nil
}
