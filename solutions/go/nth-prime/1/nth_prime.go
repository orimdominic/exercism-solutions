package prime

import (
	"errors"
	"fmt"
	"math"
)

// Nth returns the nth prime number. An error must be returned if the nth prime number can't be calculated ('n' is equal or less than zero)
func Nth(n int) (int, error) {
	if n == 0 {
		return 0, errors.New("there is no zeroth prime")
	}

	if n == 1 {
		return 2, nil
	}

	currInt := 3
	nth := 1
	for currInt < math.MaxInt {
		if isPrime(currInt) {
			nth++
		}

		if nth == n {
			return currInt, nil
		}

		currInt++
	}

	return 0, fmt.Errorf("could not find prime number %d", n)
}

func isPrime(n int) bool {
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}

	return true
}
