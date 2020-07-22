package main

import (
	"math"
)

// FrogJmp returns the minimal number of jumps
func FrogJmp(X int, Y int, D int) int {
	res := (float64(Y) - float64(X)) / float64(D)
	return int(math.Ceil(res))
}
