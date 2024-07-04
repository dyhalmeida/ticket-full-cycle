package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Should_CreatePartner_Correctly(t *testing.T) {
	partnerBaseURLs := map[int]string{
		1: "http://localhost:9000/api1",
		2: "http://localhost:9000/api2",
	}
	factory := NewPartnerFactory(partnerBaseURLs)
	partner1, err := factory.CreatePartner(1)
	assert.NoError(t, err)
	assert.IsType(t, &Partner1{}, partner1)
	assert.Equal(t, "http://localhost:9000/api1", partner1.(*Partner1).BaseURL)

	partner2, err := factory.CreatePartner(2)
	assert.NoError(t, err)
	assert.IsType(t, &Partner1{}, partner2)
	assert.Equal(t, "http://localhost:9000/api2", partner2.(*Partner1).BaseURL)
}

func Test_Should_Fail_When_Creating_A_Partner_That_ID_Does_Not_Exist(t *testing.T) {
	partnerBaseURLs := map[int]string{
		1: "http://localhost:9000/api1",
	}
	factory := NewPartnerFactory(partnerBaseURLs)
	_, err := factory.CreatePartner(2)

	assert.Error(t, err)
	assert.Equal(t, "unsupported partner ID: 2", err.Error())
}
