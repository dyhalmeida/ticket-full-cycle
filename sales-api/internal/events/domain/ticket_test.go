package domain

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_Should_Create_An_Ticket_Correctly(t *testing.T) {
	event, _ := NewEvent("Concert", "Stadium", "Music Inc.", RatingFree, time.Now().Add(24*time.Hour), 100, 50.0, "http://x.jpg", 1)
	spot, _ := NewSpot(event, "A1")
	ticket, err := NewTicket(event, spot, TicketTypeFull)
	assert.Nil(t, err)
	assert.NotNil(t, ticket)
	assert.Equal(t, TicketTypeFull, ticket.TicketType)
	assert.Equal(t, 50.0, ticket.Price)
	assert.Equal(t, event.ID, ticket.EventID)
	assert.Equal(t, spot.ID, ticket.Spot.ID)
	assert.NotEmpty(t, ticket.ID)
}

func Test_Should_Create_An_Ticket_HalfPrice_Correctly(t *testing.T) {
	event, _ := NewEvent("Concert", "Stadium", "Music Inc.", RatingFree, time.Now().Add(24*time.Hour), 100, 50.0, "http://x.jpg", 1)
	spot, _ := NewSpot(event, "A1")
	ticket, err := NewTicket(event, spot, TicketTypeHalf)
	assert.Nil(t, err)
	assert.NotNil(t, ticket)
	assert.Equal(t, TicketTypeHalf, ticket.TicketType)
	assert.Equal(t, 25.0, ticket.Price)
}

func Test_Should_Create_An_Ticket_FullfPrice_Correctly(t *testing.T) {
	event, _ := NewEvent("Concert", "Stadium", "Music Inc.", RatingFree, time.Now().Add(24*time.Hour), 100, 50.0, "http://x.jpg", 1)
	spot, _ := NewSpot(event, "A1")
	ticket, err := NewTicket(event, spot, TicketTypeFull)
	assert.Nil(t, err)
	assert.NotNil(t, ticket)
	assert.Equal(t, TicketTypeFull, ticket.TicketType)
	assert.Equal(t, 50.0, ticket.Price)
}

func Test_Should_Validate_The_Ticket_Correctly(t *testing.T) {
	ticket := &Ticket{
		Price: -10,
	}

	err := ticket.Validate()
	assert.NotNil(t, err)
	assert.Equal(t, "ticket price must be greater than zero", err.Error())
}
