package distinct

import (
	"testing"

	"github.com/magiconair/properties/assert"
)

// TestDistinct tests Distinct function
func TestDistinct(t *testing.T) {
	tt := []struct {
		description string
		a           []int
		expectedRes int
	}{
		{description: "nominal case", a: []int{2, 2, 1, 2, 3, 1}, expectedRes: 3},
		{description: "nominal case", a: []int{2}, expectedRes: 1},
	}

	for _, tc := range tt {
		t.Run(tc.description, func(t *testing.T) {
			res := Distinct(tc.a)
			assert.Equal(t, res, tc.expectedRes)
		})
	}
}
