package permcheck

// PermCheck returns 1 if the array is a permutation
func PermCheck(A []int) int {

	n := len(A)
	totalSum := n * (n + 1) / 2
	var realSum int
	seen := make(map[int]bool)

	for _, e := range A {
		if seen[e] {
			// return in case of duplicate value
			return 0
		}
		seen[e] = true
		realSum += e
	}
	if totalSum == realSum {
		return 1
	}
	return 0
}
