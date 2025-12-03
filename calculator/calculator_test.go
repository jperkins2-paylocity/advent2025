package calculator

import (
	"fmt"
	"math"
	"testing"
)

// TestAdd demonstrates basic testing
func TestAdd(t *testing.T) {
	result := Add(2, 3)
	expected := 5
	if result != expected {
		t.Errorf("Add(2, 3) = %d; expected %d", result, expected)
	}
}

// TestSubtract demonstrates basic testing
func TestSubtract(t *testing.T) {
	result := Subtract(5, 3)
	expected := 2
	if result != expected {
		t.Errorf("Subtract(5, 3) = %d; expected %d", result, expected)
	}
}

// TestMultiply demonstrates table-driven tests (Go best practice)
func TestMultiply(t *testing.T) {
	tests := []struct {
		name     string
		a, b     int
		expected int
	}{
		{"positive numbers", 3, 4, 12},
		{"with zero", 5, 0, 0},
		{"negative numbers", -2, 3, -6},
		{"both negative", -2, -3, 6},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Multiply(tt.a, tt.b)
			if result != tt.expected {
				t.Errorf("Multiply(%d, %d) = %d; expected %d", tt.a, tt.b, result, tt.expected)
			}
		})
	}
}

// TestDivide demonstrates testing functions that return errors
func TestDivide(t *testing.T) {
	tests := []struct {
		name        string
		a, b        int
		expected    float64
		expectError bool
	}{
		{"normal division", 10, 2, 5.0, false},
		{"division with remainder", 7, 2, 3.5, false},
		{"division by zero", 5, 0, 0, true},
		{"negative numbers", -10, 2, -5.0, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := Divide(tt.a, tt.b)

			if tt.expectError {
				if err == nil {
					t.Errorf("Divide(%d, %d) expected error, got nil", tt.a, tt.b)
				}
			} else {
				if err != nil {
					t.Errorf("Divide(%d, %d) unexpected error: %v", tt.a, tt.b, err)
				}
				if result != tt.expected {
					t.Errorf("Divide(%d, %d) = %f; expected %f", tt.a, tt.b, result, tt.expected)
				}
			}
		})
	}
}

// TestPower demonstrates testing with floating-point comparisons
func TestPower(t *testing.T) {
	tests := []struct {
		name     string
		base     int
		exponent int
		expected float64
	}{
		{"2^3", 2, 3, 8.0},
		{"5^2", 5, 2, 25.0},
		{"3^0", 3, 0, 1.0},
		{"2^-1", 2, -1, 0.5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Power(tt.base, tt.exponent)
			if math.Abs(result-tt.expected) > 0.0001 {
				t.Errorf("Power(%d, %d) = %f; expected %f", tt.base, tt.exponent, result, tt.expected)
			}
		})
	}
}

// TestIsEven demonstrates boolean testing
func TestIsEven(t *testing.T) {
	tests := []struct {
		name     string
		input    int
		expected bool
	}{
		{"even positive", 4, true},
		{"odd positive", 5, false},
		{"zero", 0, true},
		{"even negative", -4, true},
		{"odd negative", -5, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsEven(tt.input)
			if result != tt.expected {
				t.Errorf("IsEven(%d) = %t; expected %t", tt.input, result, tt.expected)
			}
		})
	}
}

// TestFactorial demonstrates testing edge cases and errors
func TestFactorial(t *testing.T) {
	tests := []struct {
		name        string
		input       int
		expected    int
		expectError bool
	}{
		{"factorial of 0", 0, 1, false},
		{"factorial of 1", 1, 1, false},
		{"factorial of 5", 5, 120, false},
		{"factorial of 7", 7, 5040, false},
		{"negative number", -1, 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := Factorial(tt.input)

			if tt.expectError {
				if err == nil {
					t.Errorf("Factorial(%d) expected error, got nil", tt.input)
				}
			} else {
				if err != nil {
					t.Errorf("Factorial(%d) unexpected error: %v", tt.input, err)
				}
				if result != tt.expected {
					t.Errorf("Factorial(%d) = %d; expected %d", tt.input, result, tt.expected)
				}
			}
		})
	}
}

// TestSum demonstrates testing variadic functions
func TestSum(t *testing.T) {
	tests := []struct {
		name     string
		numbers  []int
		expected int
	}{
		{"empty slice", []int{}, 0},
		{"single number", []int{5}, 5},
		{"multiple numbers", []int{1, 2, 3, 4, 5}, 15},
		{"with negatives", []int{-5, 10, -3}, 2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Sum(tt.numbers...)
			if result != tt.expected {
				t.Errorf("Sum(%v) = %d; expected %d", tt.numbers, result, tt.expected)
			}
		})
	}
}

// BenchmarkAdd demonstrates benchmark testing
func BenchmarkAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Add(100, 200)
	}
}

// BenchmarkMultiply demonstrates benchmark testing
func BenchmarkMultiply(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Multiply(12, 34)
	}
}

// BenchmarkFactorial demonstrates benchmark testing with table-driven approach
func BenchmarkFactorial(b *testing.B) {
	tests := []struct {
		name  string
		input int
	}{
		{"small", 5},
		{"medium", 10},
		{"large", 15},
	}

	for _, tt := range tests {
		b.Run(tt.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				Factorial(tt.input)
			}
		})
	}
}

// ExampleAdd demonstrates example-based testing (shows up in documentation)
func ExampleAdd() {
	result := Add(2, 3)
	fmt.Println(result)
	// Output: 5
}

// ExampleDivide demonstrates example with error handling
func ExampleDivide() {
	result, err := Divide(10, 2)
	if err != nil {
		fmt.Println("Error:", err.Error())
		return
	}
	fmt.Println(result)
	// Output: 5
}
