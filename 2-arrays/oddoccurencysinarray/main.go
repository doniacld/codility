package main

// OddOccurencesInArray returns the odd value in an array
// only one value present has not pair
func OddOccurencesInArray(A []int) int {
	var oddOne int
	var presentElems = make(map[int]int)

	for _, value := range A {
		presentElems[value]++
	}

	for k, value := range presentElems {
		if isOdd(value) {
			oddOne = k
			return oddOne
		}
	}

	return -1
}
// isOdd returns the
func isOdd(n int) bool {
	return !(n%2 == 0)
}
