package core_test

import (
	"testing"

	"github.com/bmizerany/assert"
	"github.com/goboilerplates/micro-rest/core"
)

// TestNameLower .
func TestNameLower(test *testing.T) {
	sampleEntity := core.SampleEntity{Name: "AAAA"}

	assert.NotEqual(test, "AAAA", sampleEntity.NameLower())
	assert.Equal(test, "aaaa", sampleEntity.NameLower())
}
