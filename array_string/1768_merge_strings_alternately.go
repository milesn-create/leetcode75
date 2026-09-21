package array_string

import (
	"strings"
)

func mergeAlternately(word1 string, word2 string) string {
	var sb strings.Builder
	l1 := len(word1)
	l2 := len(word2)
	sb.Grow(l1 + l2)
	maxLen := max(l1, l2)
	for i := 0; i < maxLen; i++ {
		if i < l1 {
			sb.WriteByte(word1[i])
		}
		if i < l2 {
			sb.WriteByte(word2[i])
		}
	}
	return sb.String()
}
