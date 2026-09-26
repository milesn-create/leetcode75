package sliding_window

func MaxVowels(s string, k int) int {
	maxV := 0
	countV := 0
	for i := 0; i < k; i++ {
		if isVowel(s[i]) {
			countV++
		}
	}
	maxV = countV
	for i := k; i < len(s); i++ {
		if isVowel(s[i]) {
			countV++
		}
		if isVowel(s[i-k]) {
			countV--
		}
		maxV = max(maxV, countV)
	}
	return maxV
}
func isVowel(c byte) bool {
	switch c {
	case 'a', 'e', 'i', 'o', 'u':
		return true
	}
	return false
}
