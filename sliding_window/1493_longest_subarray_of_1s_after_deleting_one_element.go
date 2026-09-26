package sliding_window

func LongestSubarray(nums []int) int {
	left := 0
	maxLen := 0
	countZ := 0
	for right := 0; right < len(nums); right++ {
		if nums[right] == 0 {
			countZ++
		}
		for countZ > 1 {
			if nums[left] == 0 {
				countZ--
			}
			left++
		}
		maxLen = max(maxLen, right-left)
	}
	return maxLen
}
