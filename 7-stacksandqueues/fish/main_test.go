package fish

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFish(t *testing.T) {
	tt := []struct {
		description string
		a           []int
		b           []int
		expectedRes int
	}{
		{description: "nominal case", a: []int{4, 3, 2, 1, 5}, b: []int{0, 1, 0, 0, 0}, expectedRes: 2},
		{description: "nominal case: the first is down stream", a: []int{4, 3, 2, 1, 5}, b: []int{1, 1, 0, 0, 0}, expectedRes: 2},
		{description: "error case: 0 element in A", a: []int{}, b: []int{1, 1, 0, 0, 0}, expectedRes: -1},
		{description: "error case: 0 element in B", a: []int{0, 2}, b: []int{}, expectedRes: -1},
		{description: "error case: 0 element in B", a: []int{}, b: []int{}, expectedRes: 0},
		{description: "nominal case: only one fish", a: []int{1}, b: []int{1}, expectedRes: 1},
		{description: "nominal case: only upstream fishes", a: []int{1, 4, 5}, b: []int{0, 0, 0}, expectedRes: 3},
		{description: "nominal case: only downstream fishes", a: []int{1, 4, 5}, b: []int{1, 1, 1}, expectedRes: 3},
		{description: "nominal case: the last is downstream", a: []int{1, 6, 5}, b: []int{0, 0, 1}, expectedRes: 3},
		{description: "nominal case: the last is downstream", a: []int{1, 6, 5, 8, 7, 5, 6, 9, 8, 7, 4, 2, 5, 6, 7}, b: []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0}, expectedRes: 13},
	}

	for _, tc := range tt {
		t.Run(tc.description, func(t *testing.T) {
			res := Fish(tc.a, tc.b)
			assert.Equal(t, tc.expectedRes, res)
		})
	}
}
