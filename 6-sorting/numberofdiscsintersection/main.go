package numberofdiscsintersection

import (
	"fmt"
	"sort"
)

// NumberOfDiscIntersections returns the number of disc intersections
func NumberOfDiscIntersections(A []int) int {
	// store in two arrays the start and the end of the disks
	start := make([]int, len(A))
	end := make([]int, len(A))
	for i := range A {
		start[i] = i - A[i]
		end[i] = A[i] + i
	}
	fmt.Printf("start: %v, end: %v \n", start, end)
	// sort arrays, we do not manipulate disks but values on segments
	sort.Ints(start)
	sort.Ints(end)
	fmt.Printf("start sorted: %v, end sorted: %v \n", start, end)

	var position, intersections, activeDisks int
	for i := range start {
		if start[i] > end[position] {
			activeDisks--
			position++
		}
		if start[i] <= end[position] {
			intersections += activeDisks
			activeDisks++
			fmt.Println(i, activeDisks, intersections)
		}
		fmt.Printf("i: %v, position: %v, activeDisks: %v, intersections: %v \n", i, position, activeDisks, intersections)
	}
	return intersections
}
