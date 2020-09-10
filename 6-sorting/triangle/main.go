package triangle

import (
	"sort"
)

// Triangle returns 1 if there is at least one triangle otherwise return 0
func Triangle(A []int) int {
	if len(A) < 3 {
		// early exit
		return -1
	}
	sort.Ints(A)
	for i := 0; i < len(A)-2; i++ {
		if checkConditions(A, i, i+1, i+2) {
			return 1
		}
	}
	return 0
}

// checkConditions returns true if all the conditions are respected
func checkConditions(a []int, p, q, r int) bool {
	check1 := a[p]+a[q] > a[r]
	check2 := a[q]+a[r] > a[p]
	check3 := a[r]+a[p] > a[q]

	if check1 && check2 && check3 {
		return true
	}
	return false
}
