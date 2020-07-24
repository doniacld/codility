package mix

import "sort"

// MaxPairedOpposite returns the max value that has a paired opposite
func MaxPairedOpposite(A []int) int {
	// exit directly if the length is equal to one
	if len(A) == 1 {
		return 0
	}

	// store negative and positive values separately
	seenPositive := make(map[int]bool)
	seenNegative := make([]int, 0)
	for _, a := range A {
		if a > 0 {
			seenPositive[a] = true
		} else {
			seenNegative = append(seenNegative, a)
		}
	}

	// sort the array in order to begin by the greatest value
	sort.Ints(seenNegative)

	// check if the negative value has its opposite in the map of positive values
	for _, e := range seenNegative {
		if seenPositive[-e] {
			return -e
		}
	}

	return 0
}
