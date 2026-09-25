package two_pointers

func MaxArea(height []int) int {
	left, right := 0, len(height)-1
	maxWater := 0
	for left < right {
		maxWater = max(maxWater, (right-left)*min(height[left], height[right]))
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return maxWater
}
