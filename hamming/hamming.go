package hamming

import "fmt"

const testVersion = 5

func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return -1, fmt.Errorf("String length %d, %d not equal", len(a), len(b))
	} else {
		var count int
		for i := 0; i < len(a); i++ {
			if a[i] != b[i] {
				count++
			}
		}
		return count, nil
	}

}
