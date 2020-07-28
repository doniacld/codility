package maxproductofthree

import (
	"testing"

	"github.com/magiconair/properties/assert"
)

// TestMaxProductOfThree tests MaxProductOfThree function
func TestMaxProductOfThree(t *testing.T) {
	tt := []struct {
		description string
		a           []int
		expectedRes int
	}{
		{description: "nominal case: positive and negative values", a: []int{-5, 5, -5, 4}, expectedRes: 125},
		{description: "nominal case: only negative values", a: []int{-5, -3, -4, -2, -1}, expectedRes: -6},
		{description: "error case: length of the array is less than 3", a: []int{-5}, expectedRes: -1},
	}

	for _, tc := range tt {
		t.Run(tc.description, func(t *testing.T) {
			res := MaxProductOfThree(tc.a)
			assert.Equal(t, res, tc.expectedRes)
		})
	}
}
