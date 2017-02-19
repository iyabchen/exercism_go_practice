package sieve

// Use the Sieve of Eratosthenes to find all the primes from 2 up to a given number.
func Sieve(limit int64) []int64 {
	var ret []int64

	notprime := make([]bool, limit+1)
	var i int64
	for i = 2; i <= limit; i++ {
		if !notprime[i] {
			ret = append(ret, i)
			for j := i * 2; j <= limit; j += i {
				notprime[j] = true
			}
		}
	}
	return ret

}
