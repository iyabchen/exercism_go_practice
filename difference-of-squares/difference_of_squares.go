package diffsquares

const testVersion = 1

// an alternative could be using a formula to calculate
// eg. square of sums = (n(n+1)/2)^2
func SquareOfSums(n int) int {
	sum := 0
	for i := 1; i <= n; i++ {
		sum += i
	}
	return sum * sum
}

func SumOfSquares(n int) int {
	sum := 0
	for i := 1; i <= n; i++ {
		sum += (i * i)
	}
	return sum
}

func Difference(n int) int {
	return SquareOfSums(n) - SumOfSquares(n)

}
