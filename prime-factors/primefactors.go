package prime

import ()

const testVersion = 2

// Compute the prime factors of a given natural number.
// Return prime factors in increasing order
// eg. Given 60, return 2, 2, 3, 5

// Very slow: BenchmarkPrimeFactors-4   	     100	  11950683 ns/op
func Factors(number int64) []int64 {
	ret := []int64{}
	var i int64
	for i = 2; i <= number; {
		if number%i == 0 {
			ret = append(ret, i)
			number = number / i
		} else {
			i++
		}

	}
	return ret
}
Sieve