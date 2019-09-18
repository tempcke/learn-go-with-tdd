package string_calculator

import (
	"gotest.tools/assert"
	"testing"
)

func TestEmptyStringReturnsZero(t *testing.T) {
	input := ""
	result := resolve(input)
	assert.Equal(t, 0, result)
}