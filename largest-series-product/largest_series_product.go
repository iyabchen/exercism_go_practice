package lsproduct

import "errors"
import "unicode"

const testVersion = 4

// Given a string of digits, calculate the largest product for a contiguous substring of digits of length n.
// For example, for the input `'1027839564'`, the largest product for a
// series of 3 digits is 270 (9 * 5 * 6), and the largest product for a
// series of 5 digits is 7560 (7 * 8 * 3 * 9 * 5).
func LargestSeriesProduct(digits string, span int) (int64, error) {
	if span == 0 { // got this from the test case
		return 1, nil
	}
	strlen := len(digits)
	if span > strlen {
		return -1, errors.New("span larger than the string length")
	}
	if span < 0 {
		return -1, errors.New("span cannot be a negative number")
	}
	for _, v := range digits {
		if !unicode.IsNumber(v) {
			return -1, errors.New("string contains non number character")
		}
	}

	var max int64 = 0
	for i := 0; i <= strlen-span; i++ {
		numbers := digits[i : i+span]
		var product int64 = 1
		for j := 0; j < span; j++ {
			product = product * int64(numbers[j]-'0')
		}
		if product > max {
			max = product
		}
	}
	return max, nil
}
