package domain

import (
	"errors"
	"fmt"
)

var (
	ErrInvalidQuantity = errors.New("quantity must be greater than zero")
)

type spotService struct{}

func NewSpotService() *spotService {
	return &spotService{}
}

func (s *spotService) GenerateSpots(event *Event, quantity int) error {
	if quantity <= 0 {
		return ErrInvalidQuantity
	}

	for i := 0; i < quantity; i++ {
		spotName := fmt.Sprintf("%c%d", 'A'+i/10, i%10+1) // Gera o nome do spot como A1, A2, ..., B1, B2, ...
		spot, err := NewSpot(event, spotName)
		if err != nil {
			return err
		}
		event.Spots = append(event.Spots, *spot)
	}

	return nil
}
