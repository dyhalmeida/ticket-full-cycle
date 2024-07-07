package usecase

import (
	"github.com/dyhalmeida/ticket-full-cycle/sales-api/internal/events/domain"
)

type CreateEventUseCase struct {
	eventRepository domain.IEventRepository
}

func NewCreateEventUseCase(eventRepository domain.IEventRepository) *CreateEventUseCase {
	return &CreateEventUseCase{eventRepository: eventRepository}
}

func (uc *CreateEventUseCase) Execute(input CreateEventInputDTO) (CreateEventOutputDTO, error) {
	event, err := domain.NewEvent(
		input.Name,
		input.Location,
		input.Organization,
		domain.Rating(input.Rating),
		input.Date,
		input.Capacity,
		input.Price,
		input.ImageURL,
		input.PartnerID,
	)
	if err != nil {
		return CreateEventOutputDTO{}, err
	}

	err = uc.eventRepository.CreateEvent(event)
	if err != nil {
		return CreateEventOutputDTO{}, err
	}

	output := CreateEventOutputDTO{
		ID:           event.ID,
		Name:         event.Name,
		Location:     event.Location,
		Organization: event.Organization,
		Rating:       string(event.Rating),
		Date:         event.Date,
		Capacity:     event.Capacity,
		ImageURL:     event.ImageURL,
		Price:        event.Price,
		PartnerID:    event.PartnerID,
	}

	return output, nil
}
