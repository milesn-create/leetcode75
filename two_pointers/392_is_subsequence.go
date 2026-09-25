package two_pointers

func IsSubsequence(s string, t string) bool {
	i, j := 0, 0
	for j < len(t) && i < len(s) {
		if t[j] == s[i] {
			i++
		}
		j++
	}
	return i == len(s)
}
