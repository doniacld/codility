package main

import (
	"testing"

	"github.com/magiconair/properties/assert"
)

// TestRetrieveGapBinary tests RetrieveGapBinary function
func TestRetrieveGapBinary(t *testing.T) {
	tt := []struct {
		description string
		value       int
		expectedRes int
	}{
		{"nominal case", 5, 1},                                      // 101
		{"no gap, only ones", 15, 0},                                // 1111
		{"several gaps", 9, 2},                                      // 1001
		{"max is the first gap", 1041, 5},                           // 10000010001
		{"gap are equals", 147, 2},                                  // 10010011
		{"only one gap with several ones at the beginning", 483, 3}, // 111100011
		{"max gap is the second gap", 647, 4},                       // 1010000111
		{"max gap is the second gap", 10564, 3},                     // 10100101000100
		{"max gap is the second gap", 32, 0},                        // 10100101000100
	}

	for _, tc := range tt {
		t.Run(tc.description, func(t *testing.T) {
			res := RetrieveGapBinary(tc.value)
			assert.Equal(t, res, tc.expectedRes)
		})
	}
}
