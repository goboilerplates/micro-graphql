package api_test

import (
	"testing"

	"github.com/bmizerany/assert"
	"github.com/goboilerplates/micro-graphql/api"
)

// TestGetSampleType .
func TestGetSampleType(test *testing.T) {
	object := api.GetSampleType()
	assert.NotEqual(test, object, nil)
}
