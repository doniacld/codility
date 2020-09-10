package main

import "fmt"

func main() {
	//	fmt.Println(count(5, []int{1, 2, 3}, 5))
	fmt.Println(optimizeCount([]int{1, 3, 5}, 11))
}

/*
// count(S[], m, n)
// S[] = different solutions
// m = index of the last score
// n = total given score
// ==> count(S[], m-1, n) : solution you take the previous value
//==> count(S[], m, n-S[m]) : solution you do not choose
// Time complexity = 2^n
func count(n int, scores []int, endIdx int) int {
	if n == 0 {
		return 1
	}

	if n < 0 || endIdx < 0 {
		return 0
	}

	if scores[endIdx] > n {
		return count(n, scores, endIdx+1)
	} else {
		return count(n-scores[endIdx], scores, endIdx) + count(n, scores, endIdx-1)
	}
}
*/

// optimizeCount
func optimizeCount(scores []int, amount int) int {
	memo := make([]int, amount+1)
	memo[0] = 1
	// fill the rest of the table
	for i := 0; i < len(scores); i++ {
		for j := scores[i]; j <= amount; j++ {
			// count all the solution including the coin
			memo[j] = memo[j] + memo[j-scores[i]]
		}
	}
	return memo[amount]
}
