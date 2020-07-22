package frogriverone

// FrogRiverOne
func FrogRiverOne(X int, A []int) int {
	arraySum := 0
	presentElems := make(map[int]bool)

	expectedSum := X * (X + 1) / 2

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