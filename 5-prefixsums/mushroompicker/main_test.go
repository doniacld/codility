package mushroompicker

import (
	"testing"

	"github.com/magiconair/properties/assert"
)

func TestMushroomPicker(t *testing.T) {
	tt := []struct {
		description string
		a           []int
		k           int
		m           int
		expectedRes int
	}{
		{description: "nominal case", a: []int{2, 3, 7, 5, 1, 3, 9}, k: 4, m: 6, expectedRes: 25},
	}

	for _, tc := range tt {
		t.Run(tc.description, func(t *testing.T) {
			res := MushroomPicker(tc.a, tc.k, tc.m)
			assert.Equal(t, res, tc.expectedRes)
		})
	}
}
