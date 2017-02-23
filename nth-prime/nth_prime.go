package prime

import "math"

const testVersion = 1

// Given a number n, determine what the nth prime is.
// eg.
// By listing the first six prime numbers: 2, 3, 5, 7, 11, and 13, we can see that
// the 6th prime is 13.

// global variable, each time findPrimes is invoked, it will add some elements if found
var primes = []int{}

// Find primes in the start - end range, and add to array primes. If the number
// of primes reaches the target, then break early
// The correctness of this function relies on primes correctly storing all primes
// in [2,start).
// This function uses sieve algorithm
// The purpose of the function using a range is map[int]bool cannot be set to
//  a very large capacity or else, there is allocation failure
func findPrimes(start int, end int, target int) {

	nonprimes := make(map[int]bool, end-start+1)

	// mark non primes from existing primes
	var p int // mark down the number at the end of prime array
	for _, p = range primes {
		var i int
		if p*p >= start {
			i = p * p
		} else if start%p == 0 {
			i = start
		} else {
			i = (start/p + 1) * p
		}
		for ; i <= end; i = i + p {
			nonprimes[i] = true
		}
	}

	// start adding new primes
	sqrt_end := int(math.Sqrt(float64(end)))
	var i int
	if p < start {
		i = start
	} else { // p >=start, to address prime number 2
		i = p + 1
	}
	for ; i <= sqrt_end; i++ { // use i*i <=end is also OK, if sqrt result not required
		if !nonprimes[i] {
			primes = append(primes, i)
			if len(primes) == target {
				return
			}
			// mark non primes
			for j := i * i; j <= end; j = j + i {
				nonprimes[j] = true
			}
		}
	}
	// from sqrt(end) to end, the unmarked are primes
	// start might be larger than sqrt_end in previous step
	if sqrt_end+1 < start {
		i = start
	} else {
		i = sqrt_end + 1
	}
	for ; i <= end; i++ {
		if !nonprimes[i] {
			primes = append(primes, i)
			if len(primes) == target {
				return
			}
		}
	}

}

// Return the nth prime number, and whether there is error
func Nth(n int) (int, bool) {
	if n == 0 {
		return 0, false
	}

	inc := math.MaxInt16
	for start, end := 2, inc; len(primes) < n; start, end = end+1, end+inc {
		findPrimes(start, end, n)
	}
	return primes[n-1], true
}
