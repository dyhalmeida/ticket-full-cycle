package usecase

import "github.com/dyhalmeida/ticket-full-cycle/sales-api/internal/events/domain"

type GetEventUseCase struct {
	eventRepository domain.IEventRepository
}

func NewGetEventUseCase(eventRepository domain.IEventRepository) *GetEventUseCase {
	return &GetEventUseCase{eventRepository: eventRepository}
}

func (uc *GetEventUseCase) Execute(input GetEventInputDTO) (*GetEventOutputDTO, error) {
	event, err := uc.eventRepository.FindEventByID(input.ID)
	if err != nil {
		return nil, err
	}

	return &GetEventOutputDTO{
		ID:           event.ID,
		Name:         event.Name,
		Location:     event.Location,
		Organization: event.Organization,
		Rating:       string(event.Rating),
		Date:         event.Date.Format("2006-01-02 15:04:05"),
		Capacity:     event.Capacity,
		Price:        event.Price,
		PartnerID:    event.PartnerID,
	}, nil
}
