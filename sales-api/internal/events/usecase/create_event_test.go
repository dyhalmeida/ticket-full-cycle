package usecase

import (
	"testing"
	"time"

	"github.com/dyhalmeida/ticket-full-cycle/sales-api/internal/events/domain"
	"github.com/dyhalmeida/ticket-full-cycle/sales-api/internal/events/infra/repository"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_Create_Event_UseCase_Execute(t *testing.T) {
	mockRepo := new(repository.MockEventRepository)
	createEventUseCase := NewCreateEventUseCase(mockRepo)

	input := CreateEventInputDTO{
		Name:         "Concert",
		Location:     "Stadium",
		Organization: "Music Inc.",
		Rating:       string(domain.RatingFree),
		Date:         time.Now().Add(24 * time.Hour),
		Capacity:     100,
		Price:        50.0,
		PartnerID:    1,
	}

	mockRepo.On("CreateEvent", mock.AnythingOfType("*domain.Event")).Return(nil)

	output, err := createEventUseCase.Execute(input)

	assert.Nil(t, err)
	assert.NotNil(t, output)
	assert.Equal(t, input.Name, output.Name)
	assert.Equal(t, input.Location, output.Location)
	assert.Equal(t, input.Organization, output.Organization)
	assert.Equal(t, input.Rating, output.Rating)
	assert.WithinDuration(t, input.Date, output.Date, time.Second)
	assert.Equal(t, input.Capacity, output.Capacity)
	assert.Equal(t, input.Price, output.Price)
	assert.Equal(t, input.PartnerID, output.PartnerID)

	mockRepo.AssertExpectations(t)
}
