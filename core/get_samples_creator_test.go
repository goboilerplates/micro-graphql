package core_test

import (
	"testing"

	"github.com/goboilerplates/micro-rest/core"
	"github.com/stretchr/testify/assert"
)

// TestCreateDefaultGetSamples .
func TestCreateDefaultGetSamples(test *testing.T) {
	interactor := core.CreateDefaultGetSamples()
	assert.NotNil(test, interactor)
}
