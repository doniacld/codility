package frogriverone

// FrogRiverOne return the earliest time when the frog
// can jump to the other side of the river
func FrogRiverOne(X int, A []int) int {
	arraySum := 0
	presentElems := make(map[int]bool)

	// compute the sum using the function Sum(n) = n*(n+1)/2
	expectedSum := X * (X + 1) / 2

	// browse all values until all the elements have been seen
	for i, position := range A {
		if !presentElems[position] {
			arraySum += position
			presentElems[position] = true
		}

		if arraySum == expectedSum {
			return i
		}
	}
	return -1
}