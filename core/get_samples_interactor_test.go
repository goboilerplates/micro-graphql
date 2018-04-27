package core_test

import (
	"testing"

	"github.com/goboilerplates/micro-rest/core"
	"github.com/stretchr/testify/assert"
)

func SetupGetSamplesInteractor() core.GetSamplesInteractor {
	return core.CreateDefaultGetSamples("Kaka", "Ronaldo")
}

// TestGetSamplesAll .
func TestGetSamplesAll(test *testing.T) {
	interactor := SetupGetSamplesInteractor()

	samples, err := interactor.All()
	assert.Nil(test, err)
	assert.Equal(test, len(samples), 2)
}

// TestGetSamplesAllByName .
func TestGetSamplesAllByName(test *testing.T) {
	interactor := SetupGetSamplesInteractor()

	samples, err := interactor.AllByName("ka")
	assert.Nil(test, err)
	assert.Equal(test, len(samples), 1)

	samples, err = interactor.AllByName("b")
	assert.NotNil(test, err)
	assert.Equal(test, len(samples), 0)

	samples, err = interactor.AllByName("a")
	assert.Nil(test, err)
	assert.Equal(test, len(samples), 2)

	samples, err = interactor.AllByName("nal")
	assert.Nil(test, err)
	assert.Equal(test, len(samples), 1)
}
