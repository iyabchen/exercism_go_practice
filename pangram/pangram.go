package pangram

import "strings"

const testVersion = 1

func IsPangram(s string) bool {
	var arr [26]bool

	for _, ch := range strings.ToLower(s) {
		// unicode.isLetter is not good to use to judge alphabets
		// ch is int32/rune type
		if 97 <= ch && 123 >= ch {
			arr[ch-'a'] = true
		}
	}

	for _, b := range arr {
		if b == false {
			return false
		}

	}
	return true

}
