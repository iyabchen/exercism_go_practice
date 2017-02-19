package palindrome

import (
	"errors"
)

const testVersion = 1
const MaxUint = ^uint(0)
const MinUint = 0
const MaxInt = int(MaxUint >> 1)
const MinInt = -MaxInt - 1

// Given the range `[1, 9]` (both inclusive)...

// The smallest palindromic product is `1`. It's factors are `(1, 1)`.
// The largest palindromic product is `9`. It's factors are `(1, 9)`, `(3, 3)`. `(9, 1)`
// is repeated, and should not be returned.

type Product struct {
	// palindromic number
	Product int
	//list of all possible two-factor factorizations of Product, within given limits, in order
	Factorizations [][2]int
}

func Products(fmin, fmax int) (Product, Product, error) {
	if fmin > fmax {
		return Product{}, Product{}, errors.New("fmin > fmax")
	}
	minProduct, maxProduct := Product{MaxInt, nil}, Product{MinInt, nil}
	for i := fmin; i <= fmax; i++ {
		for j := i; j <= fmax; j++ {

			if isPalindromic(i * j) {
				if i*j < minProduct.Product {
					minProduct = Product{i * j, [][2]int{{i, j}}}
				} else if i*j == minProduct.Product {
					minProduct.Factorizations = append(minProduct.Factorizations, [2]int{i, j})
				}
				if i*j > maxProduct.Product {
					maxProduct = Product{i * j, [][2]int{{i, j}}}
				} else if i*j == maxProduct.Product {
					maxProduct.Factorizations = append(maxProduct.Factorizations, [2]int{i, j})
				}
			}
		}
	}
	if minProduct.Factorizations == nil {
		return Product{}, Product{}, errors.New("No palindromes")
	}
	return minProduct, maxProduct, nil

}

func isPalindromic(n int) bool {
	n1 := n
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
	return rev == n1

}
