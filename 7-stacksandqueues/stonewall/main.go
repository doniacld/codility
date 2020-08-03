package stonewall

// StoneWall returns the number of stone wall needed to complete the building
func StoneWall(H []int) int {

	if len(H) == 0 {
		return 0
	}

	stack := make([]int, 0)
	var countBlock int

	for i := range H {
		for len(stack) > 0 && stack[len(stack)-1] > H[i] {
			// remove blocks bigger than the given height
			stack = append(stack[:len(stack)-1]) // stack.Pop()
		}
		if len(stack) != 0 && stack[len(stack)-1] == H[i] {
			// do nothing if the value is already present in the stack
			continue
		} else {
			// a new block is required, let's store it
			stack = append(stack, H[i]) // stack.Push()
			countBlock++
		}
	}

	return countBlock
}
