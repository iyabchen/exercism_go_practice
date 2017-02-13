package pascal

const testVersion = 1

func Triangle(n int) [][]int {
	if n < 1 {
		panic("The parameter n must be >= 1.")
	}

	ret := make([][]int, n)

	for i := 0; i < n; i++ {
		length := i + 1
		ret[i] = make([]int, length)
		ret[i][0] = 1
		ret[i][length-1] = 1
		for j := 1; j <= i-1; j++ {
			ret[i][j] = ret[i-1][j-1] + ret[i-1][j]
		}
	}

	return ret

}
