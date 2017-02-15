package sieve

func Sieve(limit int) []int {
	var ret []int

	var notprime []bool = make([]bool, limit+1)

	for i := 2; i <= limit; i++ {
		if !notprime[i] {
			ret = append(ret, i)
			for j := i * 2; j <= limit; j += i {
				notprime[j] = true
			}
		}
	}
	return ret

}
