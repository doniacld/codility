package maxproductofthree

import (
	"math"
	"sort"
)

// MaxProductOfThree find the maximum of a triplet in an array of size n
func MaxProductOfThree(A []int) int {
	arrayLen := len(A)
	if arrayLen < 3 {
		// early exit
		return -1
	}

	// sort the input array in an ascendant array
	sort.Ints(A)

	// computes the product of the three last elements
	// computes the product of the two first element (if two negative values, it become a positive one)
	// return the maximum product
	productTwoFirstEltsAndLast := A[0] * A[1] * A[arrayLen-1]
	productThreeLastElts := A[arrayLen-1] * A[arrayLen-2] * A[arrayLen-3]
	productMax := math.Max(float64(productTwoFirstEltsAndLast), float64(productThreeLastElts))

	return int(productMax)
}
