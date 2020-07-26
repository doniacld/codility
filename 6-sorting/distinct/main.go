package distinct

// Distinct returns the number of distinct values in an array
func Distinct(A []int) int {
	seen := make(map[int]bool)
	for _, v := range A {
		if !seen[v] {
			seen[v] = true
		}
	}

	return len(seen)
}
