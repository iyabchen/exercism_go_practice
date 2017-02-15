package luhn

import "strings"
import "unicode"

const testVersion = 1

func Valid(src string) bool {
	if len(src) <= 1 {
		return false
	}

	str := StripSpace(src)
	for _, v := range str {
		if !unicode.IsNumber(v) {
			return false
		}
	}

	cnt := 0
	sum := 0
	for i := len(str) - 1; i >= 0; i-- {
		cnt++
		n := int(str[i] - '0')
		if cnt%2 == 0 {
			if n*2 > 9 {
				n = n*2 - 9
			} else {
				n = n * 2
			}
		}
		sum += n
	}
	if sum%10 == 0 {
		return true
	} else {
		return false
	}

}

func StripSpace(s string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsSpace(r) {
			return -1
		}
		return r
	}, s)
}
