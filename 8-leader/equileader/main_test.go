package equileader

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestEquileader tests Equileader function
func TestEquileader(t *testing.T) {
	tt := []struct {
		description string
		a           []int
		expectedRes int
	}{
		{description: "nominal case", a: []int{6, 2, 3, 6, 6, 6, 6, 8, 1}, expectedRes: 6},
	}

	for _, tc := range tt {
		t.Run(tc.description, func(t *testing.T) {
			res := Equileader(tc.a)
			assert.Equal(t, tc.expectedRes, res)
		})
	}
}
