package main

import (
	"fmt"

	"github.com/jperkins2-paylocity/advent2025/calculator"
	"github.com/jperkins2-paylocity/advent2025/stringutil"
)

func main() {
	fmt.Println("=== Calculator Examples ===")
	
	// Basic arithmetic
	fmt.Printf("Add(5, 3) = %d\n", calculator.Add(5, 3))
	fmt.Printf("Subtract(10, 4) = %d\n", calculator.Subtract(10, 4))
	fmt.Printf("Multiply(6, 7) = %d\n", calculator.Multiply(6, 7))
	
	// Division with error handling
	result, err := calculator.Divide(20, 4)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("Divide(20, 4) = %.2f\n", result)
	}
	
	// Power
	fmt.Printf("Power(2, 8) = %.0f\n", calculator.Power(2, 8))
	
	// Factorial
	fact, _ := calculator.Factorial(5)
	fmt.Printf("Factorial(5) = %d\n", fact)
	
	// Sum with variadic arguments
	fmt.Printf("Sum(1, 2, 3, 4, 5) = %d\n", calculator.Sum(1, 2, 3, 4, 5))
	
	fmt.Println("\n=== String Utilities Examples ===")
	
	// Reverse
	fmt.Printf("Reverse('hello') = '%s'\n", stringutil.Reverse("hello"))
	
	// Palindrome checking
	word := "racecar"
	fmt.Printf("IsPalindrome('%s') = %t\n", word, stringutil.IsPalindrome(word))
	
	phrase := "A man a plan a canal Panama"
	fmt.Printf("IsPalindrome('%s') = %t\n", phrase, stringutil.IsPalindrome(phrase))
	
	// Count vowels
	text := "Hello World"
	fmt.Printf("CountVowels('%s') = %d\n", text, stringutil.CountVowels(text))
	
	// Title case
	fmt.Printf("ToTitle('hello world') = '%s'\n", stringutil.ToTitle("hello world"))
	
	// Word count
	sentence := "The quick brown fox jumps over the lazy dog"
	fmt.Printf("WordCount('%s') = %d\n", sentence, stringutil.WordCount(sentence))
}
