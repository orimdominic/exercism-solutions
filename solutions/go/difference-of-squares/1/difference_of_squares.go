package diffsquares

func SquareOfSum(n int) int {
	sum := 0

	for n >= 1 {
		sum = sum + n
		n = n - 1
	}

	return sum * sum
}

func SumOfSquares(n int) int {
	return (n * (n+1) * ((2 * n )+ 1))/6
}

func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
