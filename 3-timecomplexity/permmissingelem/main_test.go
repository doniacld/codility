package permmissingelem

import (
	"testing"

	"github.com/magiconair/properties/assert"
)

// TestPermMissingElem tests PermMissingElem function
func TestPermMissingElem(t *testing.T) {

	tt := []struct {
		description string
		input       []int
		expectedRes int
	}{
		{description: "zero element", input: []int{}, expectedRes: 1},
		{description: "one element", input: []int{2}, expectedRes: 1},
		{description: "missing 4", input: []int{2, 3, 1, 5}, expectedRes: 4},
		{description: "missing 10", input: []int{2, 3, 1, 4, 5, 7, 9, 6, 8, 11}, expectedRes: 10},
	}

	for _, tc := range tt {
		t.Run(tc.description, func(t *testing.T) {
			res := PermMissingElem(tc.input)
			assert.Equal(t, res, tc.expectedRes)
		})
	}
}
