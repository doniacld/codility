package maxcounters

// MaxCounters returns counters incremented :
// if A[K] = X, such that 1 ≤ X ≤ N, then operation K is increase(X),
// if A[K] = N + 1 then operation K is max counter.
func MaxCounters(N int, A []int) []int {
	counters := make([]int, N)
	var currentMaxCounter, floorCounter int

	for _, value := range A {
		// set the floor counter to the current max counter
		if value == N+1 {
			floorCounter = currentMaxCounter
		} else {
			// update the current value with the floor counter
			if counters[value-1] < floorCounter {
				counters[value-1] = floorCounter
			}

			counters[value-1]++

			// update current max counter
			if currentMaxCounter < counters[value-1] {
				currentMaxCounter = counters[value-1]
			}
		}
	}
	// update all the values under the floor counter
	for i, value := range counters {
		if value < floorCounter {
			counters[i] = floorCounter
		}
	}
	return counters
}
