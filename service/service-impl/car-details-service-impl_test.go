package service_impl

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	sut = NewCarDetailsService()
)

func TestGetDetails(t *testing.T) {
	carDetails := sut.GetDetails()

	assert.NotNil(t, carDetails)
	assert.Equal(t, 1, carDetails.Id)
	assert.Equal(t, "Mitsubishi", carDetails.Brand)
	assert.Equal(t, 2002, carDetails.Year)
}
