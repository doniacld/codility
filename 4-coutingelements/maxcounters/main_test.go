package maxcounters

import (
	"testing"

	"github.com/magiconair/properties/assert"
)

// TestMaxCounters tests MaxCounters function
func TestMaxCounters(t *testing.T) {
	tt := []struct {
		description string
		x           int
		a           []int
		expectedRes []int
	}{
		{description: "nominal case", x: 5, a: []int{3, 4, 4, 6, 1, 4, 4}, expectedRes: []int{3, 2, 2, 4, 2}},
	}

	for _, tc := range tt {
		t.Run(tc.description, func(t *testing.T) {
			res := MaxCounters(tc.x, tc.a)
			assert.Equal(t, res, tc.expectedRes)
		})
	}
}
