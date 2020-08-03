package mix

import (
	"strconv"
	"strings"
)

//  n! , (n-1)!,..., 1 en O(n)
func Factorials(n int) string {
	var result []int
	for i := 1; i < n+1; i++ {
		result = append(result, factorial(i))
	}
	reverseResult := reverseSlice(result)
	res := strings.Join(reverseResult, ", ")
	return res
}

func factorial(n int) int {
	if n == 0 {
		return 1
	} else {
		return n * factorial(n-1)
	}
}

func reverseSlice(a []int) []string {
	reversedArray := make([]string, 0)
	for i := len(a) - 1; i >= 0; i-- {
		reversedArray = append(reversedArray, strconv.Itoa(a[i]))
	}
	return reversedArray
}
