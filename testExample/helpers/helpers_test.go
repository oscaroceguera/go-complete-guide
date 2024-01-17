package helpers

import (
	"testing"

	"github.com/magiconair/properties/assert"
)

func TestSum(t *testing.T) {
	h := NewHelpers()
	result := h.Sum(1,2)

	assert.Equal(t, 3, result)
	
}

func TestSumMultiples(t *testing.T) {
	h := NewHelpers()
	numbers := []int{1,2,1,1}

	result := h.SumMultiple(numbers)

	assert.Equal(t, 5, result)
	
}
