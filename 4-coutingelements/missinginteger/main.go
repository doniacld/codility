package missinginteger

// MissingInteger returns the smallest positive integer that does not occur in a given array
func MissingInteger(A []int) int {
	seen := make(map[int]bool)
	// fill the map with present values
	for _, value := range A {
		if value > 0 && !seen[value] {
			seen[value] = true
		}
	}

	// browse values until the missing one is found
	missingInt := 1
	for i := missingInt; i <= len(seen); i++ {
		if !seen[missingInt] {
			return missingInt
		}
		missingInt++
	}
	return missingInt
}
