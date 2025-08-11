package armstrong

import (
	"math"
)

// 1234
func IsNumber(n int) bool {
	ord := orderOfDigits(n)
	temp := n
	sum := 0

	for i := 0; i < ord; i++ {
		mod := temp % 10
		sum = sum + int(math.Pow(float64(mod), float64(ord)))
		temp = int(temp / 10)
	}
	return sum == n
}

// orderOfDigits returns the number of digits in an int n
func orderOfDigits(n int) int {
	if n == 0 {
		return 1
	}

	ord := int(math.Log10(math.Abs(float64(n)))) + 1
	return int(ord)
}
