package hash_map_set

func FindDifference(nums1 []int, nums2 []int) [][]int {
	set1 := make(map[int]bool)
	set2 := make(map[int]bool)
	var result1, result2 []int
	for _, n := range nums1 {
		set1[n] = true
	}
	for _, n := range nums2 {
		set2[n] = true
	}
	for n := range set1 {
		if !set2[n] {
			result1 = append(result1, n)
		}
	}
	for n := range set2 {
		if !set1[n] {
			result2 = append(result2, n)
		}
	}
	return [][]int{result1, result2}
}
