package mushroompicker

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// TestMushroomPicker tests MushroomPicker function
func TestMushroomPicker(t *testing.T) {
	tt := []struct {
		description string
		a           []int
		k           int
		m           int
		expectedRes int
	}{
		{description: "nominal case", a: []int{2, 3, 7, 5, 1, 3, 9}, k: 4, m: 6, expectedRes: 25},
		//{description: "invalid case: empty array", a: []int{}, k: 4, m: 6, expectedRes: -1},
	}

	for _, tc := range tt {
		t.Run(tc.description, func(t *testing.T) {
			res := MushroomPicker(tc.a, tc.k, tc.m)
			assert.Equal(t, tc.expectedRes, res)
		})
	}
}
