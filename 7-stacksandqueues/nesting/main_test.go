package nesting

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestNesting tests Nesting function
func TestNesting(t *testing.T) {
	tt := []struct {
		description string
		a           string
		expectedRes int
	}{
		{description: "nominal case", a: "((()()))", expectedRes: 1},
		{description: "error case: number of elements are pair", a: "(()(", expectedRes: 0},
		{description: "error case: number of elements are odd", a: "(()", expectedRes: 0},
		{description: "nominal case no elements", a: "", expectedRes: 1},
		{description: "error case: more closing bracket at the end", a: "(()))", expectedRes: 0},
		{description: "error case: begin by a closing bracket", a: ")(((", expectedRes: 0},
		{description: "error case: only opening brackets", a: "((((", expectedRes: 0},
	}

	for _, tc := range tt {
		t.Run(tc.description, func(t *testing.T) {
			res := Nesting(tc.a)
			assert.Equal(t, tc.expectedRes, res)
		})
	}
}
