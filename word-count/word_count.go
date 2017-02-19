package wordcount

import "strings"
import "regexp"

const testVersion = 3

// Use this return type.
type Frequency map[string]int

// Count the frequency of words in a phrase, ignoring case
// the phrase could contain forms in different case, separate by
// space or comma or semicolon or else, words could be quoted
// eg. don't is a word, but 'abc' need to be stripped as abc
// For more details check the test cases
func WordCount(phrase string) Frequency {
	var f Frequency = make(Frequency)
	phrase = strings.ToLower(phrase)

	re := regexp.MustCompile("[a-z0-9]+('[a-z0-9]+|[a-z0-9]*)")
	words := re.FindAllString(phrase, -1)

	for _, v := range words {
		if len(v) > 0 {
			f[v]++
		}

	}
	return f
}
