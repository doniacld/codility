package distinct

// Distinct returns the number of distinct values in an array
func Distinct(A []int) int {
	if len(A) == 0 {
		// early exit
		return 0
	}

	seen := make(map[int]bool)
	for _, v := range A {
		if !seen[v] {
			seen[v] = true
		}
	}

	return len(seen)
}

/* Another solution without maps
func Distinct(A []int) int {
	arrayLen := len(A)
	sort.Ints(A)
	result := 1

	if arrayLen == 0 {
		// early exit
		return 0
	}

	for i := 1; i < arrayLen; i += 1 {
		if A[i] != A[i - 1] {
			result += 1
		}
	}

	return result
}
 */
