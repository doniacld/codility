package main

import (
	"math"
	"strconv"
	"strings"
)

// RetrieveGapBinary retrieves any maximal sequence of consecutive zeros that is surrounded by ones
func RetrieveGapBinary(N int) int {
	// convert the given integer into its binary form
	valueB := strconv.FormatInt(int64(N), 2)

	// needed values during the browse of the binary value
	// declared as float64 to avoid casting in the loop
	var gapLength, maxGapLength float64

	// browse the binary value and count the number of consecutive zeros
	for _, valueS := range strings.Split(valueB, "") {
		if valueS == "0" {
			gapLength++
		} else {
			// the value one is encountered
			// either it is the end of a sequence or there is no sequence of zeros
			// the max gap length is retrieved
			maxGapLength = math.Max(maxGapLength, gapLength)
			gapLength = 0
		}
	}

	return int(maxGapLength)
}
