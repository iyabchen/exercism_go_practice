package grains

import "errors"

const testVersion = 1
const size = 64

func Square(n int) (uint64, error) {
	if n > size || n < 1 {
		return 0, errors.New("Out of range, input must be [1-64]")
	}
	return 1 << uint64(n-1), nil

}

func Total() uint64 {
	var ret uint64 = 0
	return ^ret

}
