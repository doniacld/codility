package main

import (
	"fmt"
)

/*
You are given coins of different denominations and a total amount of money amount. Write a function to compute the fewest number of coins that you need to make up that amount. If that amount of money cannot be made up by any combination of the coins, return -1.

Example 1:

Input: coins = [1, 2, 5], amount = 11
Output: 3
Explanation: 11 = 5 + 5 + 1

Example 2:

Input: coins = [2], amount = 3
Output: -1

Note:
You may assume that you have an infinite number of each kind of coin.
*/

func main() {
	fmt.Println(coinChange([]int{1, 3, 5}, 11))
}

func coinChange(coins []int, money int) int {
	// make a DP array
	numCoins := len(coins)
	dp := make([]int, money+1)
	dp[0] = 1
	for i := 0; i < numCoins; i++ {
		for j := coins[i]; j <= money; j++ {
			// use the saved solution
			dp[j] += dp[j-coins[i]]
		}
	}
	return dp[money]
}
/*
// coinChanges
func coinChange(coins []int, amount int) int {
	// early exit
	if amount == 0 {
		return 0
	}

	// create the dynamic array that we will need
	nbCoins := make([]int, amount+1)

	// we do not need the first value, let's just set it
	nbCoins[0] = 0
	for _, coin := range coins {
		if coin > amount {
			continue
		}
		nbCoins[coin] = 1
		for i := coin + 1; i <= amount; i++ {
			if nbCoins[i-coin] == 0 {
				continue
			}
			if nbCoins[i] == 0 || nbCoins[i-coin]+1 < nbCoins[i] {
				nbCoins[i] = nbCoins[i-coin] + 1
			}
		}
	}
	if nbCoins[amount] == 0 {
		return -1
	}

	return nbCoins[amount]
}
*/
