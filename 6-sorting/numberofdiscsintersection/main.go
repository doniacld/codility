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
	// sort arrays, we do not manipulate disks but values on segments
	fmt.Println(start, end)
	sort.Ints(start)
	sort.Ints(end)
	fmt.Println(start, end)

	var position int
	for i := range start {
		// return the index of the start array
		// start[i] <= end[i]
		position = sort.SearchInts(end, start[i])
		fmt.Println(position)

	}
	return 0
}
