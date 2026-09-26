package prefix_sum

func largestAltitude(gain []int) int {
	maxA := 0
	curA := 0
	for i := 0; i < len(gain); i++ {
		curA = curA + gain[i]
		maxA = max(maxA, curA)

	}
	return maxA

}
