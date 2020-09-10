package main

import (
	"fmt"
	"math"
)

// https://www.youtube.com/watch?v=jgiZlGzXMBw
func main() {
	fmt.Println(count([]int{1, 2, 5}, 11))
}

func count(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := range dp {
		dp[i] = amount + 1
	}
	dp[0] = 0
	for i := 1; i < len(dp); i++ {
		for j := 0; j < len(coins); j++ {
			if i-coins[j] < 0 {
				continue
			}
			dp[i] = int(math.Min(float64(dp[i]), float64(1+dp[i-coins[j]])))
		}
	}
	return dp[amount]
}
