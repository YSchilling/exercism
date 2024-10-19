package thefarm

import (
	"errors"
	"fmt"
)

// TODO: define the 'DivideFood' function
func DivideFood(calculator FodderCalculator, nCows int) (float64, error) {
	totalFodder, err := calculator.FodderAmount(nCows)
	if err != nil {
		return 0, err
	}

	fatteningFactor, err := calculator.FatteningFactor()
	if err != nil {
		return 0, err
	}

	return totalFodder * fatteningFactor / float64(nCows), nil
}

// TODO: define the 'ValidateInputAndDivideFood' function
func ValidateInputAndDivideFood(calculator FodderCalculator, nCows int) (float64, error) {
	if nCows <= 0 {
		return 0, errors.New("invalid number of cows")
	}

	return DivideFood(calculator, nCows)
}

// TODO: define the 'ValidateNumberOfCows' function
type InvalidCowsError struct {
	nCows   int
	message string
}

func (p *InvalidCowsError) Error() string {
	return fmt.Sprintf("%d cows are invalid: %s", p.nCows, p.message)
}

func ValidateNumberOfCows(nCows int) error {
	if nCows < 0 {
		return &InvalidCowsError{nCows, "there are no negative cows"}
	}

	if nCows == 0 {
		return &InvalidCowsError{nCows, "no cows don't need food"}
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
