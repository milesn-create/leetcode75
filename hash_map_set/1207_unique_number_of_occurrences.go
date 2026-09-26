package hash_map_set

func uniqueOccurrences(arr []int) bool {
	count := make(map[int]int)

	for _, n := range arr {
		count[n]++
	}
	seen := make(map[int]bool)
	for _, c := range count {
		if seen[c] == true {
			return false
		}
		seen[c] = true

	}
	return true

}
