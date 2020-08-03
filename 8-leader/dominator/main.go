package dominator

// Dominator returns the value of the dominator in a given array
// dominator means that the value that appears the most is present more than half the array
func Dominator(A []int) int {

	if len(A) == 0 {
		// early exit
		return -1
	}
	if len(A) == 1 {
		// early exit
		return 0
	}

	n := len(A)
	var size, value int
	for i := range A {
		// the stack is empty, let's fill it
		if size == 0 {
			size++
			value = A[i]
		} else { // otherwise check the current value with the the stackValue
			if value != A[i] {
				size--
			} else { // in this case store the value
				size++
			}
		}
	}

	// let's store the candidate position
	var candidate int
	if size > 0 {
		candidate = value
	}

	// let's count the number of times it appears
	var count, position int
	for i := range A {
		if A[i] == candidate {
			count++
			position = i
		}
	}
	// if it appears more than half of the length, return the candidate
	if count > n/2 {
		return position
	}
	return -1
}
