package main

import "fmt"

func main() {
	res := permute([]int{1, 2, 3})
	fmt.Println(res)
}

func permute(nums []int) [][]int {

	result := [][]int{{}}
	// browse given nums, compute the combination for each value of the array
	for _, value := range nums {

		currentResult := [][]int{}

		// browse result
		for _, perm := range result {
			// browse permutation array
			for i := 0; i < len(perm)+1; i++ {
				item := make([]int, len(perm)+1)
				// store the current value we want to move
				item[i] = value
				// create the combination
				copy(item[:i], perm[:i]) // end of the array
				copy(item[i+1:], perm[i:]) // beginning of the array

				currentResult = append(currentResult, item)
			}
		}
		result = currentResult
	}
	return result
}