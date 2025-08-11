package grains

import (
	"errors"
	"math"
)

func Square(number int) (uint64, error) {
	// geometric progression a*r^(n-1)
	if number < 1 || number > 64 {
		return 0, errors.New("invalid number")
	}
	commonRatio := 2
	return uint64(math.Pow(float64(commonRatio), float64(number-1))), nil
}

func Total() uint64 {
	// geometric sum
	s := 1
	for i := 0; i < 64; i++ {
		s = s * (2)
	}

	return uint64(s - 1)
}
