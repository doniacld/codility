package triangle

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestTriangle tests Triangle function
func TestTriangle(t *testing.T) {
	tt := []struct {
		description string
		a           []int
		expectedRes int
	}{
		{description: "nominal case: 1 triangle", a: []int{10, 2, 5, 1, 8, 20}, expectedRes: 1},
		{description: "no triangle", a: []int{10, 50, 5, 1}, expectedRes: 0},
		{description: "no triangle: negative value", a: []int{-10, 50, 5, 1}, expectedRes: 0},
		{description: "no triangle: 0 elements", a: []int{}, expectedRes: -1},
		{description: "no triangle: 2 elements", a: []int{0, 1}, expectedRes: -1},
	}

	for _, tc := range tt {
		t.Run(tc.description, func(t *testing.T) {
			res := Triangle(tc.a)
			assert.Equal(t, tc.expectedRes, res)
		})
	}
}
