package main

// CyclicRotation returns the given array rotated to the right K times
func CyclicRotation(A []int, K int) []int {
	length := len(A)

	if length == 0 {
		// early exit
		return A
	}

	// Number of rotations is superior to the length,
	// let's divide K by the length.
	if K > length {
		K = K % length
	}
	// take the array from the position where the rotation begin until the end
	// and add the the beginning part
	A = append(A[length-K:length], A[0:length-K]...)

	return A
}
