package scrabble

import "strings"
import "unicode"

const testVersion = 4

var mapping = map[int][]string{
	1:  {"A", "E", "I", "O", "U", "L", "N", "R", "S", "T"},
	2:  {"D", "G"},
	3:  {"B", "C", "M", "P"},
	4:  {"F", "H", "V", "W", "Y"},
	5:  {"K"},
	8:  {"J", "X"},
	10: {"Q", "Z"},
}

func Score(s string) int {
	score_table := Transform(mapping)
	score := 0
	for _, v := range strings.ToLower(s) {
		if unicode.IsLetter(v) {
			score += score_table[string(v)]
		}
	}
	return score
}

func Transform(src map[int][]string) map[string]int {
	ret := make(map[string]int, 26)
	for key, value := range src {
		for _, letter := range value {
			ret[strings.ToLower(letter)] = key
		}
	}
	return ret
}
