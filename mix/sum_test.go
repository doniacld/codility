package mix

import (
	"testing"

	"github.com/magiconair/properties/assert"
)


// TestSum tests Sum function
func TestSum(t *testing.T) {
	tt := []struct {
		description string
		a           []int
		expectedRes int
	}{
		{description: "nominal case", a: []int{2, 3, 4, 1, 5}, expectedRes: 2},
		{description: "nominal case", a: []int{1, 3, 4, 2, 5}, expectedRes: 3},
		{description: "weird case with no suite", a: []int{1, 3, 6}, expectedRes: 1},
		{description: "only one element which is not one", a: []int{5}, expectedRes: 0},
		{description: "only one element that is one", a: []int{1}, expectedRes: 1},
	}

	for _, tc := range tt {
		t.Run(tc.description, func(t *testing.T) {
			res := Sum(tc.a)
			assert.Equal(t, res, tc.expectedRes)
		})
	}
}
