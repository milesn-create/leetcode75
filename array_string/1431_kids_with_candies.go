package array_string

import "slices"

func KidsWithCandies(candies []int, extraCandies int) []bool {
	result := make([]bool, len(candies))
	maxCandies := slices.Max(candies)
	for i := range candies {
		if candies[i]+extraCandies > maxCandies {
			result[i] = true
		} else {
			result[i] = false
		}
	}

	return result
}
