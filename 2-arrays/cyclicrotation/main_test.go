package main

import (
	"testing"

	"github.com/magiconair/properties/assert"
)

// TestCyclicRotation tests TestCyclicRotation function
func TestCyclicRotation(t *testing.T) {
	tt := []struct {
		description string
		A           []int
		k           int
		expectedRes []int
	}{
		{"no rotation", []int{3, 8, 9, 7, 6}, 0, []int{3, 8, 9, 7, 6}},
		{"one rotation", []int{3, 8, 9, 7, 6}, 1, []int{6, 3, 8, 9, 7}},
		{"two rotations", []int{3, 8, 9, 7, 6}, 2, []int{7, 6, 3, 8, 9}},
		{"three rotations", []int{3, 8, 9, 7, 6}, 3, []int{9, 7, 6, 3, 8}},
		{"two rotations with negative number", []int{-10, 8, 9, 7, 6}, 2, []int{7, 6, -10, 8, 9}},
		{"two rotations with the same numbers", []int{0, 0, 0, 0}, 2, []int{0, 0, 0, 0}},
		{"no rotation because length and K are equals", []int{3, 8, 9, 7, 6}, 5, []int{3, 8, 9, 7, 6}},
	}

	for _, tc := range tt {
		t.Run(tc.description, func(t *testing.T) {
			res := CyclicRotation(tc.A, tc.k)
			assert.Equal(t, res, tc.expectedRes)
		})
	}
}
