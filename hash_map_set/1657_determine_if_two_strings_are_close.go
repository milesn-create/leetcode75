package hash_map_set

func СloseStrings(word1 string, word2 string) bool {
	if len(word1) != len(word2) {
		return false
	}
	c1 := make(map[rune]int)
	c2 := make(map[rune]int)
	for _, ch := range word1 {
		c1[ch] += 1
	}
	for _, ch := range word2 {
		c2[ch] += 1
	}
	if len(c1) != len(c2) {
		return false
	}
	for ch := range c1 {
		if _, ok := c2[ch]; !ok {
			return false
		}
	}
	freq1 := make(map[int]int)
	freq2 := make(map[int]int)
	for _, f := range c1 {
		freq1[f]++
	}
	for _, f := range c2 {
		freq2[f]++
	}
	for f, count := range freq1 {
		if freq2[f] != count {
			return false
		}
	}
	return true
}
