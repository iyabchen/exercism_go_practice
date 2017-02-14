package etl

import "strings"

// given a mapping with the style of int:string, return
// a mapping with the style string:int
func Transform(src map[int][]string) map[string]int {
	ret := make(map[string]int, 26)
	for key, value := range src {
		for _, letter := range value {
			ret[strings.ToLower(letter)] = key
		}
	}
	return ret
}
