package countdiv

// CountDiv returns the number of integers within the range [A..B] that are divisible by K
func CountDiv(A, B, K int) int {
	count := B/K - A/K
	if A%K == 0 {
		count++
	}
	return count
}
