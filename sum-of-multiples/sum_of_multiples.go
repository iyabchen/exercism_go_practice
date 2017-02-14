package summultiples

import "sort"

//Given a number, find the sum of all the multiples of particular
// numbers up to but not including that number.
// divisors should not be changed, causing test result weird
func SumMultiples(limit int, divisors ...int) int {

	// remove repeated element and same mod elements
	// add multiplies of the rest of the elements

	// remove repeated and also remove divisor that larger than limit
	mapping := make(map[int]int)
	for _, v := range divisors {
		if v < limit {
			mapping[v] = 1
		}
	}
	// copy to sort
	arr := make([]int, len(mapping))
	index := 0
	for key, _ := range mapping {
		arr[index] = key
		index++
	}
	sort.Ints(arr)
	// remove same mode larger number
	for i := 0; i < len(arr); i++ {
		if arr[i] < 0 {
			continue
		}
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < 0 {
				continue
			}
			if arr[j]%arr[i] == 0 {
				delete(mapping, arr[j])
				arr[j] = -1
			} else if arr[j]*arr[i] < limit {
				mapping[arr[j]*arr[i]] = -1
			}
		}
	}

	sum := 0
	for key, value := range mapping {
		n := (limit - 1) / key
		if n == 0 {
			continue
		}
		an := key * n
		if value > 0 {
			sum += SumAP(key, an, n)
		} else {
			sum -= SumAP(key, an, n)
		}

	}
	return sum
}

func SumAP(a1, an, n int) int {
	return ((a1 + an) * n / 2)
}
