package darts

import "math"

func Score(x, y float64) int {
	// from Pythagoras' theorem
	hyp := math.Sqrt((math.Pow(x,2) +math.Pow(y, 2)))

	if hyp <= 1 {
		return 10
	}

	if hyp > 1 && hyp <= 5 {
		return 5
	}

	if hyp > 5 && hyp <= 10 {
		return 1
	}

	return 0
}
