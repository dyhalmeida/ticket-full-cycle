package usecase

import "github.com/dyhalmeida/ticket-full-cycle/sales-api/internal/events/domain"

type ListSpotsUseCase struct {
	eventRepository domain.IEventRepository
}

func NewListSpotsUseCase(eventRepository domain.IEventRepository) *ListSpotsUseCase {
	return &ListSpotsUseCase{eventRepository: eventRepository}
}

func (uc *ListSpotsUseCase) Execute(input ListSpotsInputDTO) (*ListSpotsOutputDTO, error) {
	event, err := uc.eventRepository.FindEventByID(input.EventID)
	if err != nil {
		return nil, err
	}

	spots, err := uc.eventRepository.FindSpotsByEventID(input.EventID)
	if err != nil {
		return nil, err
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

	eventDTO := EventDTO{
		ID:           event.ID,
		Name:         event.Name,
		Location:     event.Location,
		Organization: event.Organization,
		Rating:       string(event.Rating),
		Date:         event.Date.Format("2006-01-02 15:04:05"),
		Capacity:     event.Capacity,
		Price:        event.Price,
		PartnerID:    event.PartnerID,
		ImageURL:     event.ImageURL,
	}

	return &ListSpotsOutputDTO{Event: eventDTO, Spots: spotsDTO}, nil
}
