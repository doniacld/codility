package main

import (
	"testing"

	"github.com/magiconair/properties/assert"
)

// TestFrogJmp tests FrogJmp function
func TestFrogJmp(t *testing.T) {

	tt := []struct {
		description string
		x           int
		y           int
		d           int
		expectedRes int
	}{
		{"nominal case", 10, 85, 30, 3},
		{"nominal case", 10, 85, 8, 10},
	}

	for _, tc := range tt {
		t.Run(tc.description, func(t *testing.T) {
			res := FrogJmp(tc.x, tc.y, tc.d)
			assert.Equal(t, res, tc.expectedRes)
		})
	}
}
