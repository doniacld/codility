package mix

import (
	"strconv"
	"strings"
)

// Factorials print this: n! , (n-1)!,..., 1
// Time complexity O(n)
func Factorials(n int) string {
	if n == 0 {
		return ""
	}
	a := make([]int, 0)
	_, a = factorial(n, a)
	res := reverseSlice(a)
	return strings.Join(res, ", ")
}

func factorial(n int, a []int) (int, []int) {
	var res int
	if n == 0 {
		return 1, a
	} else {
		res, a = factorial(n-1, a)
		a = append(a, res*n)
		return res * n, a
	}
}

func reverseSlice(a []int) []string {
	reversedArray := make([]string, 0)
	for i := len(a) - 1; i >= 0; i-- {
		reversedArray = append(reversedArray, strconv.Itoa(a[i]))
	}
	return reversedArray
}
