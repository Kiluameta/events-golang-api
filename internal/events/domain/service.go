package domain

import (
	"errors"
	"fmt"
)

var (
	ErrInvalidQuantityGreater = errors.New("quantity must be greater than zero")
)

type SpotService struct{}

func NewSpotService() *SpotService {
	return &SpotService{}
}

func (s *SpotService) GenerateSpotService(event *Event, quantity int) error {
	if quantity <= 0 {
		return ErrInvalidQuantityGreater
	}

	for i := range quantity {
		// ADD VALIDATIONS IN THE FUTURE
		spotName := fmt.Sprintf("%c%d", 'A'+i/10, i%10+1);
		spot, err := NewSpot(event, spotName);
		if err != nil {
			return err
		}
		event.Spots = append(event.Spots, *spot);
	}

	return nil
}