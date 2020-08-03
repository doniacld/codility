package dominator

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// TestDominator tests Dominator function
func TestDominator(t *testing.T) {
	tt := []struct {
		description string
		a           []int
		expectedRes int
	}{
		{description: "nominal case", a: []int{6, 2, 3, 6, 6, 6, 6, 8, 1}, expectedRes: 6},
		{description: "nominal case: even values", a: []int{3, 4, 3, 2, 3, -1, 3, 3}, expectedRes: 7},
		{description: "nominal case: 1 element", a: []int{1}, expectedRes: 0},
		{description: "error case: no dominator", a: []int{1, 2, 3, 5, 6}, expectedRes: -1},
		{description: "error case: no dominator", a: []int{1, 1, 3, 5, 6}, expectedRes: -1},
		{description: "error case: 0 element", a: []int{}, expectedRes: -1},
	}

	for _, tc := range tt {
		t.Run(tc.description, func(t *testing.T) {
			res := Dominator(tc.a)
			assert.Equal(t, tc.expectedRes, res)
		})
	}
}
