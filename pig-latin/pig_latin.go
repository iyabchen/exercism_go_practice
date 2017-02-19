package igpay

import "strings"

// Implement a program that translates from English to Pig Latin

// Pig Latin is a made-up children's language that's intended to be
// confusing. It obeys a few simple rules (below), but when it's spoken
// quickly it's really difficult for non-children (and non-native speakers)
// to understand.

// - **Rule 1**: If a word begins with a vowel sound, add an "ay" sound to
//   the end of the word.
// - **Rule 2**: If a word begins with a consonant sound, move it to the
//   end of the word, and then add an "ay" sound to the end of the word.

// need to be familiar with english to find the underlying rules
// completely using test case to fabricate this function

// an alternative is to use map to make isVowel search easier and quicker
var vowel = []byte{'a', 'e', 'i', 'o', 'u'}

func PigLatin(src string) string {
	strArr := strings.Fields(src)
	for i, word := range strArr {
		if isVowel(word[0]) {
			strArr[i] = word + "ay"
		} else if word[0] == 'x' || word[0] == 'y' {
			if !isVowel(word[1]) { // if len(word)==1, then invalid word
				strArr[i] = word + "ay"
			} else {
				strArr[i] = word[1:] + string(word[0]) + "ay"
			}
		} else {
			inx := searchFirstVowelIndex(word)
			if inx > 0 {
				strArr[i] = word[inx:] + word[0:inx] + "ay"
			}

		}

	}
	return strings.Join(strArr, " ")

}

// search vowel starts from where
// exception, words containing qu+vowel
func searchFirstVowelIndex(word string) int {
	for i := 0; i < len(word); i++ {
		if isVowel(word[i]) {
			if i > 0 && word[i-1] == 'q' {
				return i + 1
			} else {
				return i
			}
		}
	}
	return -1

}

// find whether s is in the vowel array
func isVowel(s byte) bool {
	for _, v := range vowel {
		if s == v {
			return true
		}
	}
	return false
}
