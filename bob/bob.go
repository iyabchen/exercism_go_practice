package bob // package name must match the package name in bob_test.go

import "strings"

const testVersion = 2 // same as targetTestVersion

const (
	QUESTION = "Sure."
	YELL     = "Whoa, chill out!"
	NOWORD   = "Fine. Be that way!"
	ELSE     = "Whatever."
)

func Hey(input string) string {
	trim := strings.TrimSpace(input)
	if trim == "" {
		return NOWORD
	}
	// do lower to make sure it contains letter
	// if there is no letter, than upper and lower return the same result
	if strings.ToUpper(trim) == trim && strings.ToLower(trim) != trim {
		return YELL
	}
	if strings.HasSuffix(trim, "?") {
		return QUESTION
	}
	return ELSE
}
