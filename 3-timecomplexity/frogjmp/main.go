package main

import (
	"math"
)

// FrogJmp returns the minimal number of jumps
func FrogJmp(X int, Y int, D int) int {
	if  D <= 0 {
		// early exit
		return -1
	}

	// divide the difference of position by the distance of a step
	res := (float64(Y) - float64(X)) / float64(D)
	return int(math.Ceil(res))
}
