package wordy

import "strings"
import "strconv"

const testVersion = 1

// Parse and evaluate simple math word problems returning the answer as an integer.
// input: What is 5 plus 13?
// return 18
// if not valid question, then return false
func Answer(question string) (int, bool) {
	question = strings.TrimSuffix(question, "?")
	words := strings.Fields(question)
	numbers := extractNumbers(question)
	if len(numbers) == 0 {
		return 0, false
	}

	ret := numbers[0]
	cnt := 1
	noerr := false
	for _, word := range words {
		switch word {
		case "multiplied":
			ret = ret * numbers[cnt]
			cnt++
			noerr = true
		case "plus":
			ret = ret + numbers[cnt]
			cnt++
			noerr = true
		case "minus":
			ret = ret - numbers[cnt]
			cnt++
			noerr = true
		case "divided":
			ret = ret / numbers[cnt]
			cnt++
			noerr = true
		}
	}
	if cnt != len(numbers) {
		return 0, false
	}
	return ret, noerr
}

func extractNumbers(question string) []int {
	numbers := []int{}
	for _, word := range strings.Fields(question) {
		num, err := strconv.Atoi(word)
		if err == nil {
			numbers = append(numbers, num)
		}

	}
	return numbers
}
