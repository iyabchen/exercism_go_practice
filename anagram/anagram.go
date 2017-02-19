package anagram

import (
	"strings"
)

const testVersion = 1

// Given a word and a list of possible anagrams, select the correct sublist.
// which are anagrams of that word
// anagram is defined as a word formed from another by
// rearranging its letters, but not equal to the original word
// ignoring the case. eg. ant & nat

func Detect(word string, candidates []string) []string {
	ret := []string{}
	word = strings.ToLower(word)
	for _, v := range candidates {
		v = strings.ToLower(v)
		if isAnagram(word, v) {
			ret = append(ret, v)
		}

	}
	return ret
}

func isAnagram(word, v string) bool {

	if len(v) != len(word) || v == word {
		return false
	}

	wordMap := map[byte]int{}
	for i := 0; i < len(word); i++ {
		wordMap[word[i]]++
		wordMap[v[i]]--
	}

	for _, v := range wordMap {
		if v != 0 {
			return false
		}
	}
	return true

}
