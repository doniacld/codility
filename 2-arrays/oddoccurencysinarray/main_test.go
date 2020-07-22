package main

import (
	"testing"

	"github.com/magiconair/properties/assert"
)

// TestOddOccurencesInArray tests OddOccurencesInArray function
func TestOddOccurencesInArray(t *testing.T) {

	tt := []struct {
		description string
		array       []int
		expectedRes int
	}{
		{"nominal case", []int{1, 9, 1, 1, 5, 1, 9, 7, 5}, 7},
		{"", []int{1, 9, 10, 1, 1, 5, 1, 9, 7, 5}, 10},
	}
	for _, tc := range tt {
		t.Run(tc.description, func(t *testing.T) {
			res := OddOccurencesInArray(tc.array)
			assert.Equal(t, res, tc.expectedRes)
		})
	}
}

// TestIsOdd
func TestIsOdd(t *testing.T) {

	res := isOdd(5)
	assert.Equal(t, true, res)
}
