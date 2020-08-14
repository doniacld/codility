package stonewall

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestStoneWall tests StoneWall method
func TestStoneWall(t *testing.T) {
	tt := []struct {
		description    string
		input          []int
		expectedOutput int
	}{
		{description: "nominal case", input: []int{8, 8, 5, 7, 9, 8, 7, 4, 8}, expectedOutput: 7},
		{description: "nominal case: other ", input: []int{1, 2, 2, 3, 1}, expectedOutput: 3},
		{description: "nominal case: 0 elements", input: []int{}, expectedOutput: 0},
		{description: "nominal case: 0 elements", input: []int{5}, expectedOutput: 1},
	}

	for _, tc := range tt {
		t.Run(tc.description, func(t *testing.T) {
			output := StoneWall(tc.input)
			assert.Equal(t, tc.expectedOutput, output)
		})
	}
}
