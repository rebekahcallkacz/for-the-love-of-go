// Package calculator does simple calculations.
package calculator

import (
	"errors"
)

// Add takes two numbers and returns the result of adding
// them together.
func Add(a, b float64) float64 {
	return a + b
}

// Subtract takes two numbers a and b, and
// returns the result of subtracting b from a.
func Subtract(a, b float64) float64 {
	return a - b
	}

// Multiply takes two numbers a and b, and
// returns the result of multiplying a by b.
func Multiply(a, b  float64) float64 {
	return a * b
}

// Divide takes two numbers a and b, and
// returns the result of dividiving a by b.
func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("division by 0 is not allowed")
	}
	return a / b, nil
}