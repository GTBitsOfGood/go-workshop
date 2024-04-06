package main

import "testing"

func TestDivideNonZeroDivisor(t *testing.T) {
	quotient, err := Divide(10, 5)

	if quotient != 2 || err != nil {
		t.Fatalf(
			`Divide(10, 5) = %v, %v, expected %v, %v`,
			quotient, err,
			2, nil,
		)
	}
}

func TestDivideZeroDivisor(t *testing.T) {
	quotient, err := Divide(10, 0)

	if quotient != 0 || err != CannotDivideByZeroError {
		t.Fatalf(
			`Divide(10, 0) = %v, %v, expected %v, %v`,
			quotient, err,
			0, CannotDivideByZeroError,
		)
	}
}
