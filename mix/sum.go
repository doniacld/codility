package mix

// Sum computes the value when the Sum n*(n+1)/2
func Sum(A []int) int {
	var turnOn, realSum int
	for i := 0; i < len(A); i++ {
		expectedSum := (i + 1) * (i + 2) / 2
		realSum += A[i]
		// compare the expected sum computed with the function
		// and the real sum with given value
		if realSum == expectedSum {
			turnOn++
		}
	}

	return turnOn
}

