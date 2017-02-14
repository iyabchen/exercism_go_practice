package isogram

import "strings"
import "unicode"

const testVersion = 1

func IsIsogram(str string) bool {
	m := make(map[rune]int)
	for _, v := range strings.ToLower(str) {
		if unicode.IsLetter(v) {
			m[v]++

			if m[v] > 1 {
				return false
			}
		}
	}
	return true
}
