// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package greeting should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package greeting

// testVersion identifies the version of the test program that you are
// writing your code to. If the test program changes in the future --
// after you have posted this code to the Exercism site -- nitpickers
// will see that your code can't necessarily be expected to pass the
// current test suite because it was written to an earlier test version.
const testVersion = 3

// @param s string
// @return Hello, World! if input is empty, other wise,
// return Hello, <string>!
func HelloWorld(s string) string {
	if s == "" {
		return "Hello, World!"
	} else {
		return "Hello, " + s + "!"
	}
}
