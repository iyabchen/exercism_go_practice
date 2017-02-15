package cryptosquare

import (
	"math"
	"strings"
	"unicode"
)

const testVersion = 2

// seems able to encode, but decode is irreversible
// since normalized
func Encode(src string) string {

	// to lowercase, and remove non letter
	var normArr []rune

	for _, v := range strings.ToLower(src) {
		if unicode.IsLetter(v) || unicode.IsNumber(v) {
			normArr = append(normArr, v)
		}
	}
	normArrLen := len(normArr)

	// calculate c and r, c>=r, c-r<=1
	// r is not actually needed in transpose
	sqrt := math.Sqrt(float64(normArrLen))
	c := int(math.Ceil(sqrt))

	//	normArr - r*c, transpose the matrix and get runeArr - c*r,
	runeArr := []rune{}
	cnt := 0
	for i := 0; i < c; i++ {
		for j := 0; j < normArrLen; j += c {
			if i+j < normArrLen {
				runeArr = append(runeArr, normArr[i+j])
				cnt++
			}
		}
		if i < c-1 {
			runeArr = append(runeArr, ' ')
		}
	}

	return string(runeArr)

}
