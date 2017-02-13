// Leap stub file

// The package name is expected by the test program.
package leap

// testVersion should match the targetTestVersion in the test file.
const testVersion = 3

// given a year and return whether that year is leap year
// 1> if not century years, leap year is divisible by 4
// 2> if century years, leap year is divisible by 400
func IsLeapYear(year int) bool {
	if year%4 == 0 {
		if year%100 != 0 {
			return true
		} else if year%400 == 0 {
			return true
		}
	}
	return false
}
