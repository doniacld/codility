package mix

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountNumberWords(t *testing.T) {
	tt := map[string]struct {
		input       string
		expectedRes int
	}{
		"nominal case": {input: "I am the one. Do you think  so really?", expectedRes: 4},
	}

	for n, tc := range tt {
		t.Run(n, func(t *testing.T) {
			res := CountNumberWords(tc.input)
			assert.Equal(t, tc.expectedRes, res)
		})
	}
}
