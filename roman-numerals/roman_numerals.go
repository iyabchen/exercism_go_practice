package romannumerals

import "errors"
import "bytes"
import "strings"

const testVersion = 3

// The Romans wrote numbers using letters - I, V, X, L, C, D, M.
// I - 1 : 1-3
// V - 5 : 4-8
// X - 10 : 9-39
// L - 50 : 40-89
// C - 100: 90-399
// D - 500: 400-899
// M - 1000: 900-3999

// this style looks much shorter than the if else version,
// but does not mean faster
var translateTable = []struct {
	arabic int
	roman  string
}{
	{1, "I"}, {4, "IV"}, {5, "V"}, {9, "IX"}, {10, "X"}, {40, "XL"}, {50, "L"},
	{90, "XC"}, {100, "C"}, {400, "CD"}, {500, "D"}, {900, "CM"}, {1000, "M"},
}

// Given a number n, convert it into a roman number
// 0<=n<=3000

func ToRomanNumeral(n int) (string, error) {
	if n <= 0 || n >= 4000 {
		return "", errors.New("n must be within range 1-3999")
	}
	var buf bytes.Buffer
	for {
		if n == 0 {
			break
		}
		for i := len(translateTable) - 1; i >= 0; i-- {
			cnt := n / translateTable[i].arabic
			n = n % translateTable[i].arabic
			buf.WriteString(strings.Repeat(translateTable[i].roman, cnt))
		}

	}
	return buf.String(), nil
}
