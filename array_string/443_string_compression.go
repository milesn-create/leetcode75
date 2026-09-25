package array_string

import "strconv"

func Сompress(chars []byte) int {
	read, write := 0, 0
	for read < len(chars) {
		ch := chars[read]
		count := 0
		for read < len(chars) && chars[read] == ch {
			read++
			count++

		}
		chars[write] = ch
		write++
		strCount := strconv.Itoa(count)
		if count > 1 {
			for i := range strCount {
				chars[write] = strCount[i]
				write++
			}
		}
	}
	return write
}
