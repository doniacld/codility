package mix

import (
	"testing"

	"github.com/magiconair/properties/assert"
)

// TestMaxPairedOpposite tests MaxPairedOpposite function
func TestMaxPairedOpposite(t *testing.T) {
	tt := []struct {
		description string
		a           []int
		expectedRes int
	}{
		{description: "3 is the biggest value with its pair", a: []int{-3, 2, 3, -2, -1, 1, 6}, expectedRes: 3},
		{description: "only one element", a: []int{1}, expectedRes: 0},
		{description: "no pair", a: []int{1, 2}, expectedRes: 0},
	}

	for _, tc := range tt {
		t.Run(tc.description, func(t *testing.T) {
			res := MaxPairedOpposite(tc.a)
			assert.Equal(t, res, tc.expectedRes)
		})
	}
}
