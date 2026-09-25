package array_string

import "math"

func IncreasingTriplet(nums []int) bool {
	first, second := math.MaxInt, math.MaxInt
	for _, n := range nums {
		if n > second {
			return true
		} else if n > first {
			second = n
		} else {
			first = n
		}
	}
	return false
}
