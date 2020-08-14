package missinginteger

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestMissingInteger tests MissingInteger function
func TestMissingInteger(t *testing.T) {
	tt := []struct {
		description string
		a           []int
		expectedRes int
	}{
		{description: "nominal case", a: []int{1, 2, 3}, expectedRes: 4},
		{description: "nominal case with a negative value", a: []int{-1}, expectedRes: 1},
		{description: "only negative numbers", a: []int{-1, -3}, expectedRes: 1},
	}

	for _, tc := range tt {
		t.Run(tc.description, func(t *testing.T) {
			res := MissingInteger(tc.a)
			assert.Equal(t, tc.expectedRes, res)
		})
	}
}
