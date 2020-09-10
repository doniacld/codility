package mushroompicker

import (
	"math"
)

// MushroomPicker returns the optimal way to pick the maximum of mushrooms
func MushroomPicker(A []int, k int, m int) int {
	if len(A) == 0 {
		return -1
	}

	// compute the prefix array
	prefix := computePrefix(A)

	// just for the readability
	move := float64(m)
	kF := float64(k)
	lenWay := float64(len(A) - 1)

	var res float64

	maxKM := int(math.Min(move, kF))
	for i := 0; i < maxKM; i++ {
		left := k - i
		right := math.Min(math.Max(float64(k+m-2*i), kF), lenWay)
		res = math.Max(res, float64(computeDiff(prefix, left, int(right))))
	}

	maxLenAandK := int(math.Min(lenWay-kF, move+1))
	for i := 0; i < maxLenAandK; i++ {
		left := math.Max(0, math.Min(float64(k-(m-2*i)), kF))
		right := k + i
		res = math.Max(res, float64(computeDiff(prefix, int(left), right)))
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

// computeDiff returns the difference between two values in an array
func computeDiff(prefix []int, start, end int) int {
	return prefix[end] - prefix[start]
}
