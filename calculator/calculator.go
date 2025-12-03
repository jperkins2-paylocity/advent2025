package calculator

import (
	"errors"
	"math"
)

// Add returns the sum of two integers
func Add(a, b int) int {
	return a + b
}

// Subtract returns the difference between two integers
func Subtract(a, b int) int {
	return a - b
}

// Multiply returns the product of two integers
func Multiply(a, b int) int {
	return a * b
}

// Divide returns the quotient of two integers
// Returns an error if the divisor is zero
func Divide(a, b int) (float64, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return float64(a) / float64(b), nil
}

// Power returns the base raised to the exponent
func Power(base, exponent int) float64 {
	return math.Pow(float64(base), float64(exponent))
}

// IsEven checks if a number is even
func IsEven(n int) bool {
	return n%2 == 0
}

// Factorial calculates the factorial of a non-negative integer
// Returns an error for negative numbers
func Factorial(n int) (int, error) {
	if n < 0 {
		return 0, errors.New("factorial is not defined for negative numbers")
	}
	if n == 0 || n == 1 {
		return 1, nil
	}
	result := 1
	for i := 2; i <= n; i++ {
		result *= i
	}
	return result, nil
}

// Sum calculates the sum of a variable number of integers
func Sum(numbers ...int) int {
	total := 0
	for _, num := range numbers {
		total += num
	}
	return total
}
