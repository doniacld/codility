package numberofdiscsintersection

import (
	"sort"
)

var maxIntersects = 10000000

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
	sort.Ints(start)
	sort.Ints(end)

	// loop until the start values are all treated
	// increment whether end index or start index
	var idxStart, idxEnd, intersections, activeDisks int
	for idxStart < len(start) {
		// We treat the current start value, we close one disk and pass to the next value of end array.
		if start[idxStart] > end[idxEnd] {
			activeDisks--
			idxEnd++
		}
		// The current disk intersects all the active ones, there is a new active disks and we can treat the next start value.
		if start[idxStart] <= end[idxEnd] {
			intersections += activeDisks
			// early exit if the max intersections is reached
			if intersections > maxIntersects {
				return -1
			}
			activeDisks++
			idxStart++
		}
	}

	return intersections
}
