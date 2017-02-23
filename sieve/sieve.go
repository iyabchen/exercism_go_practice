package sieve

import "math"

// Use the Sieve of Eratosthenes to find all the primes from 2 up to a given number.
func Sieve(limit int64) []int64 {
	var ret []int64

	notprime := make([]bool, limit+1)
	var i int64
	sqrt := int64(math.Sqrt(float64(limit)))
	for i = 2; i <= sqrt; i++ {
		if !notprime[i] {
			ret = append(ret, i)
			for j := i * i; j <= limit; j += i {
				notprime[j] = true
			}
		}
	}
	for i = sqrt + 1; i <= limit; i++ {
		if !notprime[i] {
			ret = append(ret, i)
		}
	}
	return ret

}
