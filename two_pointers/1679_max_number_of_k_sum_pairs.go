package two_pointers

import "slices"

func MaxOperations(nums []int, k int) int {
	slices.Sort(nums)
	left, right := 0, len(nums)-1
	count := 0
	for left < right {
		sum := nums[right] + nums[left]
		if sum == k {
			right--
			left++
			count++
		} else if sum < k {
			left++
		} else {
			right--
		}
	}
	return count
}
