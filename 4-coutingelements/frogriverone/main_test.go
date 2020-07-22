package frogriverone

import (
	"testing"

	"github.com/magiconair/properties/assert"
)

// TestFrogRiverOne tests FrogRiverOne function
func TestFrogRiverOne(t *testing.T) {
	tt := []struct {
		description string
		x           int
		a           []int
		expectedRes int
	}{
		{description: "nominal case", x: 5, a: []int{1, 3, 1, 4, 2, 3, 5, 4}, expectedRes: 6},
		{description: "frog does not cross", x: 9, a: []int{1, 3, 1, 4, 2, 3, 5, 4, 10, 9, 7, 8, 5, 7, 6}, expectedRes: -1},
		{description: "frog cross", x: 10, a: []int{1, 3, 1, 4, 2, 3, 5, 4, 10, 9, 7, 8, 5, 7, 6, 8, 10}, expectedRes: 14},

	}

	for _, tc := range tt {
		t.Run(tc.description, func(t *testing.T) {
			res := FrogRiverOne(tc.x, tc.a)
			assert.Equal(t, res, tc.expectedRes)
		})
	}
}
