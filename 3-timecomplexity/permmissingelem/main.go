package permmissingelem

// PermMissingElem returns the missing element in an unsorted array
func PermMissingElem(A []int) int {
	// the first element is missing
	if len(A) == 0 {
		return 1
	}

	// compute the sum using the formula Sum(len(A)+1) - Sum(A)
	n := len(A) + 1
	sum := (n + 1) * n / 2

	var arraySum int
	for _, elem := range A {
		arraySum += elem
	}

	return sum - arraySum
}
