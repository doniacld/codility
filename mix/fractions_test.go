package mix

import (
	"testing"
)

// TestFractions tests Fractions function
func TestFractions(t *testing.T) {
	tt := []struct {
		description string
		x           []int
		y           []int
		expectedRes int
	}{
		{description: "nominal case", x: []int{1, 2, 3, 4, 0}, y: []int{2, 3, 6, 8, 4}, expectedRes: 3},
		{description: "nominal case", x: []int{1}, y: []int{2}, expectedRes: 1},
	}

	for _, tc := range tt {
		t.Run(tc.description, func(t *testing.T) {
			res := Fractions(tc.x, tc.y)
			assert.Equal(t, res, tc.expectedRes)
		})
	}
}
