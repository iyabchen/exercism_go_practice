package palindrome

import (
// "errors"
)

const testVersion = 1

type Product struct {
	// palindromic number
	Product int
	//list of all possible two-factor factorizations of Product, within given limits, in order
	Factorizations [][2]int
}

func Products(fmin, fmax int) (Product, Product, error) {
	for i := fmin; i <= fmax; i++ {
		for j := fmin; j <= fmax; j++ {
			if isPalindromic(i * j) {
				Product{Product: i * j}
			}
		}
	}
	return Product{}, Product{}, nil

}

func isPalindromic(n int) bool {
	if n < 0 {
		n = -n
	}
	rev := 0
	for {
		if n != 0 {
			rev = rev * 10
			rev += n % 10
			n = n / 10
		} else {
			break
		}
	}
	return rev == n

}
