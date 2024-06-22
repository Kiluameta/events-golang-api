package domain

import "errors"

var (
	ErrInvalidSpotNumber   = errors.New("invalid spot number")
	ErrSpotNotFound        = errors.New("spot not found")
	ErrSpotAlreadyReserved = errors.New("spot already reserved")
	ErrSpotNameRequired    = errors.New("spot name is required")
	ErrSpotNameLeast       = errors.New("spot name must be at least 2 characters long")
	ErrSpotNameStartLetter = errors.New("spot name must start with a letter")
	ErrSpotNameEndNumber   = errors.New("spot name must end with a number")
)

type SpotStatus string

const (
	SpotStatusAvailable SpotStatus = "available"
	SpotStatusReserved  SpotStatus = "reserved"
)

type Spot struct {
	ID       string
	EventID  string
	Name     string
	Status   SpotStatus
	TicketID string
}

func (s *Spot) Validate() error {
	validacoes := map[bool]error{
        s.Name == "":                        ErrSpotNameRequired,
        len(s.Name) < 2:                     ErrSpotNameLeast,
        s.Name[0] < 'A' || s.Name[0] > 'Z':  ErrSpotNameStartLetter,
        s.Name[1] < '0' || s.Name[1] > '9':  ErrSpotNameEndNumber,
    }

    for condicao, err := range validacoes {
        if condicao {
            return err
        }
    }

    return nil
}
