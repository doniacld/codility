package brackets

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestBrackets tests Brackets function
func TestBrackets(t *testing.T) {
	tt := []struct {
		description string
		a           string
		expectedRes int
	}{
		{description: "nominal case", a: "{[()()]}", expectedRes: 1},
		{description: "error case", a: "[{})]", expectedRes: 0},
		{description: "nominal case no elements", a: "", expectedRes: 1},
		{description: "error case: more close bracket at the end", a: "{[()()]}]", expectedRes: 0},
		{description: "error case: only open brackets", a: "{([(({", expectedRes: 0},
	}

	for _, tc := range tt {
		t.Run(tc.description, func(t *testing.T) {
			res := Brackets(tc.a)
			assert.Equal(t, tc.expectedRes, res)
		})
	}
}

func TestIsValidPair(t *testing.T) {
	tt := []struct {
		description string
		left        string
		right       string
		expectedRes bool
	}{
		{description: "nominal case: brackets", left: bracketOpen, right: bracketClose, expectedRes: true},
		{description: "nominal case: hook brackets", left: hookOpen, right: hookClose, expectedRes: true},
		{description: "nominal case: curly brackets", left: curlyOpen, right: curlyClose, expectedRes: true},
		{description: "error case: no match", left: curlyOpen, right: bracketClose, expectedRes: false},
	}

	for _, tc := range tt {
		t.Run(tc.description, func(t *testing.T) {
			res := isValidPair(tc.left, tc.right)
			assert.Equal(t, tc.expectedRes, res)
		})
	}
}
