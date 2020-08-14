package maxcounters

import (
	"testing"

	"github.com/magiconair/properties/assert"
)

// TestMaxCounters tests MaxCounters function
func TestMaxCounters(t *testing.T) {
	tt := []struct {
		description string
		n           int
		a           []int
		expectedRes []int
	}{
		{description: "nominal case", n: 5, a: []int{3, 4, 4, 6, 1, 4, 4}, expectedRes: []int{3, 2, 2, 4, 2}},
		{description: "invalid case: x=0", n: 0, a: []int{3, 4, 4, 6, 1, 4, 4}, expectedRes: []int{}},
		{description: "invalid case: empty array", n: 0, a: []int{}, expectedRes: []int{}},
	}

	for _, tc := range tt {
		t.Run(tc.description, func(t *testing.T) {
			res := MaxCounters(tc.n, tc.a)
			assert.Equal(t, res, tc.expectedRes)
		})
	}
}
