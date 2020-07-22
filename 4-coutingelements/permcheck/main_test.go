package permcheck

import (
	"testing"

	"github.com/magiconair/properties/assert"
)

// TestPermCheck tests PermCheck function
func TestPermCheck(t *testing.T) {
	tt := []struct {
		description string
		a           []int
		expectedRes int
	}{
		{description: "nominal case", a: []int{4, 1, 3, 2}, expectedRes: 1},
		{description: "not a perm, value is present twice", a: []int{4, 1, 3, 2, 2}, expectedRes: 0},
		{description: "not a perm, missing the first", a: []int{4, 3, 2}, expectedRes: 0},
		{description: "only one element", a: []int{4}, expectedRes: 0},
	}

	for _, tc := range tt {
		t.Run(tc.description, func(t *testing.T) {
			res := PermCheck(tc.a)
			assert.Equal(t, tc.expectedRes, res)
		})
	}
}
