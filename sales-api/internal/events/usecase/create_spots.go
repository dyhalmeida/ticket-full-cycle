package usecase

import (
	"fmt"

	"github.com/dyhalmeida/ticket-full-cycle/sales-api/internal/events/domain"
)

type CreateSpotsUseCase struct {
	eventRepository domain.IEventRepository
}

func NewCreateSpotsUseCase(eventRepository domain.IEventRepository) *CreateSpotsUseCase {
	return &CreateSpotsUseCase{eventRepository: eventRepository}
}

func (uc *CreateSpotsUseCase) Execute(input CreateSpotsInputDTO) (*CreateSpotsOutputDTO, error) {
	event, err := uc.eventRepository.FindEventByID(input.EventID)
	if err != nil {
		return nil, err
	}

	spots := make([]domain.Spot, input.NumberOfSpots)
	for i := 0; i < input.NumberOfSpots; i++ {
		spotName := generateSpotName(i)
		spot, err := domain.NewSpot(event, spotName)
		if err != nil {
			return nil, err
		}
		if err := uc.eventRepository.CreateSpot(spot); err != nil {
			return nil, err
		}
		spots[i] = *spot
	}

	spotsDTO := make([]SpotDTO, len(spots))
	for i, spot := range spots {
		spotsDTO[i] = SpotDTO{
			ID:       spot.ID,
			Name:     spot.Name,
			Status:   string(spot.Status),
			TicketID: spot.TicketID,
		}
	}

	return &CreateSpotsOutputDTO{Spots: spotsDTO}, nil
}

func generateSpotName(index int) string {
	// Gera um nome de spot baseado no índice. Ex: A1, A2, ..., B1, B2, etc.
	letter := 'A' + rune(index/10)
	number := index%10 + 1
	return fmt.Sprintf("%c%d", letter, number)
}
