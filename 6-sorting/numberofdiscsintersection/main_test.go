package numberofdiscsintersection

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestNumberOfDiscIntersections tests NumberOfDiscIntersections function
func TestNumberOfDiscIntersections(t *testing.T) {
	tt := []struct {
		description string
		a           []int
		expectedRes int
	}{
		{description: "nominal case", a: []int{1, 5, 2, 1, 4, 0}, expectedRes: 11},
	}

	for _, tc := range tt {
		t.Run(tc.description, func(t *testing.T) {
			res := NumberOfDiscIntersections(tc.a)
			assert.Equal(t, res, tc.expectedRes)
		})
	}
}
