package main

import "errors"

var (
	CannotDivideByZeroError = errors.New("cannot divide by zero")
)

func Divide(dividend, divisor float64) (float64, error) {
	if divisor == 0 {
		return 0, CannotDivideByZeroError
	}

	return dividend / divisor, nil
}
