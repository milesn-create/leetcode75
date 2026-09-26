package sliding_window

func FindMaxAverage(nums []int, k int) float64 {
	maxSum := 0
	sum := 0
	for i := 0; i < k; i++ {
		sum = sum + nums[i]
	}
	maxSum = sum
	for i := k; i < len(nums); i++ {
		sum = sum + nums[i] - nums[i-k]
		maxSum = max(sum, maxSum)
	}
	return float64(maxSum) / float64(k)

}
