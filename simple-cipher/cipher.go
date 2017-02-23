// implement caesar cipher, shift cipher, and vigenere cipher
// all returned cipher text should ignore nonalphabets, and in
// lower case
package cipher

import "strings"

type Cipher interface {
	Encode(string) string
	Decode(string) string
}

type ShiftCipher []rune

// Actual encode, right shift what s specified on plain
func shiftEncode(s []rune, plain string) string {
	runeArr := []rune{}
	cnt := 0
	for _, v := range plain {
		var val rune
		if v >= 'a' && v <= 'z' {
			val = (v-'a'+s[cnt]+26)%26 + 'a'
		} else if v >= 'A' && v <= 'Z' {
			val = (v-'A'+s[cnt]+26)%26 + 'a'
		} else {
			continue
		}
		runeArr = append(runeArr, val)
		cnt++
		if cnt == len(s) {
			cnt = 0
		}
	}
	return string(runeArr)
}

// Actual decode, left shift what s specified on cipher
func shiftDecode(s []rune, cipher string) string {
	runeArr := []rune{}
	cnt := 0

	for _, v := range cipher {
		val := v
		if v >= 'a' && v <= 'z' {
			val = (v-'a'-s[cnt]+26)%26 + 'a'
		} else if v >= 'A' && v <= 'Z' {
			val = (v-'A'-s[cnt]+26)%26 + 'A'
		} else {
			continue
		}
		runeArr = append(runeArr, val)
		cnt++
		if cnt == len(s) {
			cnt = 0
		}
	}
	return string(runeArr)
}

func (c ShiftCipher) Encode(text string) string {
	return shiftEncode(c, text)
}

func (c ShiftCipher) Decode(text string) string {
	return shiftDecode(c, text)
}

// Return a Caesar cipher, caesar does right shift 3
func NewCaesar() Cipher {
	var c ShiftCipher = []rune{3}
	return c
}

// Return a shift cipher
// input is in either in 1 to 25 or -1 to -25, 0 not allowed
func NewShift(n int) Cipher {
	if (n >= 1 && n <= 25) || (n <= -1 && n >= -25) {
		var c ShiftCipher = []rune{rune(n)}
		return c
	} else {
		return nil
	}
}

// Return a vigenere cipher
// Argument for NewVigenere must consist of lower case letters a-z only.
// Values consisting entirely of the letter 'a' are disallowed.
// For invalid arguments NewVigenere returns nil.
func NewVigenere(key string) Cipher {
	onlya := strings.Repeat("a", len(key))
	if key == onlya {
		return nil
	}
	runeArr := []rune{}
	for _, v := range key {
		if v < 'a' || v > 'z' {
			return nil
		} else {
			runeArr = append(runeArr, v-'a')
		}
	}
	var c ShiftCipher = runeArr
	return c
}
