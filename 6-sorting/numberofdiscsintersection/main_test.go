package numberofdiscsintersection

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestNumberOfDiscIntersections tests NumberOfDiscIntersections function
func TestNumberOfDiscIntersections(t *testing.T) {
	tt := []struct {
		description  string
		a            []int
		maxIntersect int
		expectedRes  int
	}{
		{description: "nominal case", a: []int{1, 5, 2, 1, 4, 0}, maxIntersect: maxIntersects, expectedRes: 11},
		{description: "error case: maxIntersect is reached", a: []int{1, 5, 2, 1, 4, 0}, maxIntersect: 4, expectedRes: -1},
		{description: "error case: 0 element", a: []int{}, maxIntersect: maxIntersects, expectedRes: 0},
		{description: "nominal case with circles = points", a: []int{1, 0, 0, 4}, maxIntersect: maxIntersects, expectedRes: 4},
	}

	for _, tc := range tt {
		t.Run(tc.description, func(t *testing.T) {
			maxIntersects = tc.maxIntersect
			res := NumberOfDiscIntersections(tc.a)
			assert.Equal(t, tc.expectedRes, res)
		})
	}
}
