package fish

func Fish(A, B []int) int {

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
	var aliveFish int

	for i := range A {
		switch {
		case B[i] == 1:
			stack = append(stack, A[i])
		case B[i] == 0 && len(stack) > 0:
			if A[i] > stack[len(stack)-1] {
				stack = append(stack[:len(stack)-1])
				aliveFish++
			}
		default:
			aliveFish++
		}
	}
	return aliveFish + len(stack)
}
