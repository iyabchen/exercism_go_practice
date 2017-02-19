package perfect

import (
	"errors"
	"math"
)

// The aliquot sum is defined as the sum of the factors of a number not
// including the number itself. For example, the aliquot sum of 15 is
// (1 + 3 + 5) = 9
// Perfect: aliquot sum = number
// Abundant: aliquot sum > number
// Deficient: aliquot sum < number

const testVersion = 1

type Classification int

const (
	ClassificationPerfect Classification = iota
	ClassificationAbundant
	ClassificationDeficient
)

var ErrOnlyPositive error = errors.New("Positive number only")

// Given a number, return its classification or error
// if not natural number, then return error
func Classify(number uint64) (Classification, error) {
	if number <= 0 {
		return -1, ErrOnlyPositive
	}

	factors := []uint64{0}

	// make sure i <= number/i
	limit := uint64(math.Floor(math.Sqrt(float64(number))))
	var i uint64
	for i = 1; i <= limit && i < number; i++ {
		if number%i == 0 {
			a, b := i, number/i
			if a == b {
				factors = append(factors, a)
			} else {
				factors = append(factors, a)
				if b != number {
					factors = append(factors, b)
				}
			}

		}
	}
	var sum uint64 = 0
	for _, v := range factors {
		sum += v
	}

	if sum == number {
		return ClassificationPerfect, nil
	} else if sum > number {
		return ClassificationAbundant, nil
	} else {
		return ClassificationDeficient, nil
	}

}
