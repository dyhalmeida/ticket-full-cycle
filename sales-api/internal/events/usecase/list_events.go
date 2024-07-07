package usecase

import "github.com/dyhalmeida/ticket-full-cycle/sales-api/internal/events/domain"

type EventDTO struct {
	ID           string  `json:"id"`
	Name         string  `json:"name"`
	Location     string  `json:"location"`
	Organization string  `json:"organization"`
	Rating       string  `json:"rating"`
	Date         string  `json:"date"`
	Capacity     int     `json:"capacity"`
	Price        float64 `json:"price"`
	PartnerID    int     `json:"partner_id"`
	ImageURL     string  `json:"image_url"`
}

type ListEventsOutputDTO struct {
	Events []EventDTO `json:"events"`
}

type ListEventsUseCase struct {
	eventRepository domain.EventRepositoryInterface
}

func NewListEventsUseCase(eventRepository domain.EventRepositoryInterface) *ListEventsUseCase {
	return &ListEventsUseCase{eventRepository: eventRepository}
}

func (uc *ListEventsUseCase) Execute() (*ListEventsOutputDTO, error) {
	events, err := uc.eventRepository.ListEvents()
	if err != nil {
		return nil, err
	}

	eventsDTO := make([]EventDTO, len(events))
	for i, event := range events {
		eventsDTO[i] = EventDTO{
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
	}

	return &ListEventsOutputDTO{Events: eventsDTO}, nil
}
