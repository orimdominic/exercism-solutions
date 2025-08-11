package thefarm

import (
	"errors"
	"fmt"
)

// TODO: define the 'DivideFood' function
func DivideFood(fc FodderCalculator, nCows int) (float64, error) {
	amt, err := fc.FodderAmount(nCows)
	if err != nil {
		return 0.0, err
	}

	fac, err := fc.FatteningFactor()
	if err != nil {
		return 0.0, err
	}

	return (amt * fac) / float64(nCows), nil
}

// TODO: define the 'ValidateInputAndDivideFood' function
func ValidateInputAndDivideFood(fc FodderCalculator, nCows int) (float64, error) {
	if nCows > 0 {
		return DivideFood(fc, nCows)
	}
	return 0.0, errors.New("invalid number of cows")
}

type InvalidCowsError struct {
	message string
}

func (ice *InvalidCowsError) Error() string {
	return ice.message
}

// TODO: define the 'ValidateNumberOfCows' function
func ValidateNumberOfCows(nCows int) error {
	if nCows < 0 {
		return &InvalidCowsError{
			message: fmt.Sprintf("%d cows are invalid: there are no negative cows", nCows),
		}
	}

	if nCows == 0 {
		return &InvalidCowsError{
			message: fmt.Sprintf("%d cows are invalid: no cows don't need food", nCows),
		}
	}

	return nil
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
