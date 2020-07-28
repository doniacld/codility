package prefixsums

/*
func mushroomPicker(A []int, k int, m int) int{
	prefixSum := prefixSum(A)

	minValue := math.Min()
	for i := 0 ; i <

	return 0
}
*/
func prefixSum(a []int) []int {
	prefixSum := make([]int,0)
	prefixSum[0] = 0

	for i := 1 ; i < len(a) ; i++ {
		prefixSum = append(prefixSum, prefixSum[i-1]+prefixSum[i])
	}
	return prefixSum
}