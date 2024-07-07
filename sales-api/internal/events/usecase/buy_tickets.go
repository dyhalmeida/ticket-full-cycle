package usecase

import (
	"github.com/dyhalmeida/ticket-full-cycle/sales-api/internal/events/domain"
	"github.com/dyhalmeida/ticket-full-cycle/sales-api/internal/events/infra/service"
)

type BuyTicketsUseCase struct {
	eventRepository domain.IEventRepository
	partnerFactory  service.IPartnerFactory
}

func NewBuyTicketsUseCase(eventRepository domain.IEventRepository, partnerFactory service.IPartnerFactory) *BuyTicketsUseCase {
	return &BuyTicketsUseCase{
		eventRepository: eventRepository,
		partnerFactory:  partnerFactory,
	}
}

func (uc *BuyTicketsUseCase) Execute(input BuyTicketsInputDTO) (*BuyTicketsOutputDTO, error) {
	event, err := uc.eventRepository.FindEventByID(input.EventID)
	if err != nil {
		return nil, err
	}

	// Cria a solicitação de reserva
	req := &service.ReservationRequest{
		EventID:    input.EventID,
		Spots:      input.Spots,
		TicketType: input.TicketType,
		CardHash:   input.CardHash,
		Email:      input.Email,
	}

	// Obtém o serviço do parceiro
	partnerService, err := uc.partnerFactory.CreatePartner(event.PartnerID)
	if err != nil {
		return nil, err
	}

	// Reserva os lugares usando o serviço do parceiro
	reservationResponse, err := partnerService.MakeReservation(req)
	if err != nil {
		return nil, err
	}

	// Salva os ingressos no banco de dados
	tickets := make([]domain.Ticket, len(reservationResponse))
	for i, reservation := range reservationResponse {
		spot, err := uc.eventRepository.FindSpotByName(event.ID, reservation.Spot)
		if err != nil {
			return nil, err
		}

		ticket, err := domain.NewTicket(event, spot, domain.TicketType(input.TicketType))
		if err != nil {
			return nil, err
		}

		err = uc.eventRepository.CreateTicket(ticket)
		if err != nil {
			return nil, err
		}

		spot.Reserve(ticket.ID)
		err = uc.eventRepository.ReserveSpot(spot.ID, ticket.ID)
		if err != nil {
			return nil, err
		}

		tickets[i] = *ticket
	}

	ticketsDTO := make([]TicketDTO, len(tickets))
	for i, ticket := range tickets {
		ticketsDTO[i] = TicketDTO{
			ID:         ticket.ID,
			SpotID:     ticket.Spot.ID,
			TicketType: string(ticket.TicketType),
			Price:      ticket.Price,
		}
	}

	return &BuyTicketsOutputDTO{Tickets: ticketsDTO}, nil
}
