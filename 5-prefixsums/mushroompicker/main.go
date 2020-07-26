package mushroompicker

import (
	"fmt"
	"math"
)

func MushroomPicker(A []int, k int, m int) int {
	// compute the prefix array
	prefix := computePrefix(A)

	// just for the readability
	mF := float64(m)
	kF := float64(k)
	lF := float64(len(A) - 1)

	var res float64

	maxKM := int(math.Min(mF, kF))
	for i := 0; i < maxKM; i++ {
		left := k - i
		right := math.Min(math.Max(float64(k+m-2*i), kF), lF)
		res = math.Max(res, float64(computeDiff(prefix, left, int(right))))
		fmt.Printf("first i: %v, left: %v, right: %v, res : %v \n", i, left, right, res)
	}

	maxLenAandK := int(math.Min(float64(len(A)-k), float64(m+1)))
	for i := 0; i < maxLenAandK; i++ {
		left := math.Max(0, math.Min(float64(k-(m-2*i)), float64(k)))
		right := k + i
		res = math.Max(res, float64(computeDiff(prefix, int(left), right)))
		fmt.Printf("secon i: %v, left: %v, right: %v, res : %v \n", i, left, right, res)
	}

	return int(res)
}

// computePrefix computes the prefix between two values
func computePrefix(a []int) []int {
	prefix := make([]int, 0)
	prefix = append(prefix, a[0])
	for i := 1; i < len(a); i++ {
		prefix = append(prefix, prefix[i-1]+a[i])
	}
	return prefix
}

// computeDiff returns the difference between two value sin an array
func computeDiff(prefix []int, start, end int) int {
	return prefix[end] - prefix[start]
}
