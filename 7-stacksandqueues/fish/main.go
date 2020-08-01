package fish

// Fish returns the number of fishes that are alive,
// if a downstream fish is encountered, we add it to the stack,
// if an upstream fish is encoutered, let's compare its size to all downstream fishes,
// in the case of the upstream fish is bigger than the the downstream, let's update the stack.
func Fish(A, B []int) int {

	// early exits
	if len(A) != len(B) {
		return -1
	}
	if len(A) == 0 {
		return 0
	}
	if len(A) == 1 {
		return 1
	}

	stack := make([]int, 0)
	var deadFishes int
	for i := range A {
		switch {
		case B[i] == 1:
			stack = append(stack, A[i]) // = stack.Push(i)
		case B[i] == 0:
			for len(stack) > 0 {
				deadFishes++
				if A[i] > stack[len(stack)-1] {
					stack = append(stack[:len(stack)-1]) // = stack.Pop()
				} else {
					break
				}
			}
		}
	}
	return len(A) - deadFishes
}
