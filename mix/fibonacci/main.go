package main

import "fmt"

func main() {
	// fmt.Println(recFib(5))
	fmt.Println(optimizeFib(5))

}

// recFib is the recursive method to compute the fibonacci suite
// complexity time:
// complexity space:
func recFib(n int) int {
	if n == 1 {
		return 0
	}
	if n == 2 {
		return 1
	} else {
		return recFib(n-1) + recFib(n-2)
	}
}

// Optimize solution
var dp = map[int]int{0: 0}

// optimizeFib is the dynamic way to compute fibonacci
// complexity time:
// complexity space:
func optimizeFib(n int) int {
	if n == 1 {
		return 0
	}
	if n == 2 {
		return 1
	} else if value, ok := dp[n]; ok {
		return value
	} else {
		value := optimizeFib(n-1) + optimizeFib(n-2)
		dp[n] = value
		return value
	}

}
