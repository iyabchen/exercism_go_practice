package acronym

import "strings"
import "unicode"

const testVersion = 2

// no standard for how to create an abbreviate
// not a good puzzle, just to learn strings and unicode
func Abbreviate(src string) string {
	var s string
	s = string(src[0])
	for i := 1; i < len(src); {
		if (src[i] == ' ' || src[i] == '-') && i+1 < len(src) {
			s += string(src[i+1])
			i++
		} else if unicode.IsUpper(rune(src[i])) && unicode.IsLower(rune(src[i-1])) {
			s += string(src[i])
		}
		i++
	}
	return strings.ToUpper(s)
}
