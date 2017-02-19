package atbash

import ()

// Implementation the atbash cipher (encoder)
// The Atbash cipher is a simple substitution cipher that relies on
// transposing all the letters in the alphabet such that the resulting
// alphabet is backwards. The first letter is replaced with the last
// letter, the second with the second-last, and so on.
// eg. Plain:  abcdefghijklmnopqrstuvwxyz
//     Cipher: zyxwvutsrqponmlkjihgfedcba

const testVersion = 1

// Encode the text into atbash cipher
// Keep only alphabets, and numbers
// The trick to encode an alphabet, 'z' - v + 'a'
func Atbash(text string) string {
	runeArr := []rune{}
	for _, v := range text {
		var char rune
		if 'A' <= v && v <= 'Z' {
			char = 'z' - v + 'A'
		} else if 'a' <= v && v <= 'z' {
			char = 'z' - v + 'a'
		} else if '0' <= v && v <= '9' {
			char = v
		} else {
			continue
		}
		if len(runeArr)%6 == 5 {
			runeArr = append(runeArr, ' ')
		}
		runeArr = append(runeArr, char)

	}
	return string(runeArr)

}
