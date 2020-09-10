package countdiv

import (
	"testing"

	"github.com/magiconair/properties/assert"
)

// TestCountDiv tests CountDiv function
func TestCountDiv(t *testing.T) {
	tt := []struct {
		description string
		a           int
		b           int
		k           int
		expectedRes int
	}{
		{description: "nominal case", a: 6, b: 11, k: 2, expectedRes: 3},
		{description: "nominal case", a: 1, b: 10, k: 1, expectedRes: 10},
		{description: "nominal case", a: 10, b: 345, k: 17, expectedRes: 20},
		{description: "nominal case", a: 10, b: 10, k: 5, expectedRes: 1},
		{description: "nominal case", a: 101, b: 100561, k: 10000, expectedRes: 10},
	}

	for _, tc := range tt {
		t.Run(tc.description, func(t *testing.T) {
			res := CountDiv(tc.a, tc.b, tc.k)
			assert.Equal(t, tc.expectedRes, res)
		})
	}
}
