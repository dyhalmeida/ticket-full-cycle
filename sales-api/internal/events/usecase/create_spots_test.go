package usecase

import (
	"testing"
	"time"

	"github.com/dyhalmeida/ticket-full-cycle/sales-api/internal/events/domain"
	"github.com/dyhalmeida/ticket-full-cycle/sales-api/internal/events/infra/repository"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_Create_Spots_UseCase_Execute(t *testing.T) {
	mockRepo := new(repository.MockEventRepository)
	createSpotsUseCase := NewCreateSpotsUseCase(mockRepo)

	eventID := "1"
	eventDate := time.Now().Add(24 * time.Hour)
	mockEvent := &domain.Event{
		ID:           eventID,
		Name:         "Concert",
		Location:     "Stadium",
		Organization: "Music Inc.",
		Rating:       domain.RatingFree,
		Date:         eventDate,
		ImageURL:     "http://example.com/image.jpg",
		Capacity:     100,
		Price:        50.0,
		PartnerID:    1,
	}

	mockRepo.On("FindEventByID", eventID).Return(mockEvent, nil)
	mockRepo.On("CreateSpot", mock.AnythingOfType("*domain.Spot")).Return(nil)

	input := CreateSpotsInputDTO{
		EventID:       eventID,
		NumberOfSpots: 2,
	}

	output, err := createSpotsUseCase.Execute(input)

	assert.Nil(t, err)
	assert.NotNil(t, output)
	assert.Equal(t, input.NumberOfSpots, len(output.Spots))

	assert.Equal(t, "A1", output.Spots[0].Name)
	assert.Equal(t, "A2", output.Spots[1].Name)

	mockRepo.AssertExpectations(t)
}
