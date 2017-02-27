// Convert a number, represented as a sequence of digits in one base, to any other base.

// Implement general base conversion. Given a number in base **a**,
// represented as a sequence of digits, convert it to base **b**.

package allyourbase

import "errors"

const testVersion = 1

var (
	ErrInvalidDigit = errors.New("Digit not in the base range")
	ErrInvalidBase  = errors.New("Base < 2")
)

func ConvertToBase(inputBase uint64, inputDigits []uint64,
	outputBase uint64) (outputDigits []uint64, err error) {
	if inputBase < 2 || outputBase < 2 {
		return nil, ErrInvalidBase
	}

	var number uint64
	for _, d := range inputDigits {
		number *= inputBase
		if d >= inputBase {
			return nil, ErrInvalidDigit
		}
		number += d
	}

	for number > 0 {
		q := number / outputBase
		r := number % outputBase
		outputDigits = append([]uint64{r}, outputDigits...)
		number = q
	}
	if len(outputDigits) == 0 {
		outputDigits = []uint64{0}
	}
	return
}
