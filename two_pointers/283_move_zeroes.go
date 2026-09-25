package two_pointers

func MoveZeroes(nums []int) {
	read, write := 0, 0
	for read < len(nums) {
		if nums[read] != 0 {
			nums[write], nums[read] = nums[read], nums[write]
			read++
			write++
		} else {
			read++
		}
	}

}
