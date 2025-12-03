package stringutil

import (
	"fmt"
	"testing"
)

// TestReverse demonstrates testing string manipulation
func TestReverse(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"simple word", "hello", "olleh"},
		{"empty string", "", ""},
		{"single character", "a", "a"},
		{"with spaces", "hello world", "dlrow olleh"},
		{"unicode", "Hello, 世界", "界世 ,olleH"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Reverse(tt.input)
			if result != tt.expected {
				t.Errorf("Reverse(%q) = %q; expected %q", tt.input, result, tt.expected)
			}
		})
	}
}

// TestIsPalindrome demonstrates testing complex logic
func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{"simple palindrome", "racecar", true},
		{"not palindrome", "hello", false},
		{"single character", "a", true},
		{"empty string", "", true},
		{"with spaces", "A man a plan a canal Panama", true},
		{"case insensitive", "RaceCar", true},
		{"with punctuation", "A man, a plan, a canal: Panama", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsPalindrome(tt.input)
			if result != tt.expected {
				t.Errorf("IsPalindrome(%q) = %t; expected %t", tt.input, result, tt.expected)
			}
		})
	}
}

// TestCountVowels demonstrates counting in strings
func TestCountVowels(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected int
	}{
		{"all vowels", "aeiou", 5},
		{"no vowels", "xyz", 0},
		{"mixed", "hello", 2},
		{"uppercase", "HELLO", 2},
		{"with spaces", "hello world", 3},
		{"empty string", "", 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := CountVowels(tt.input)
			if result != tt.expected {
				t.Errorf("CountVowels(%q) = %d; expected %d", tt.input, result, tt.expected)
			}
		})
	}
}

// TestToTitle demonstrates title case conversion
func TestToTitle(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"lowercase", "hello world", "Hello World"},
		{"mixed case", "hELLo WoRLd", "Hello World"},
		{"already title", "Hello World", "Hello World"},
		{"empty", "", ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ToTitle(tt.input)
			if result != tt.expected {
				t.Errorf("ToTitle(%q) = %q; expected %q", tt.input, result, tt.expected)
			}
		})
	}
}

// TestWordCount demonstrates counting words
func TestWordCount(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected int
	}{
		{"single word", "hello", 1},
		{"multiple words", "hello world", 2},
		{"with extra spaces", "hello  world", 2},
		{"empty string", "", 0},
		{"only spaces", "   ", 0},
		{"sentence", "The quick brown fox", 4},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := WordCount(tt.input)
			if result != tt.expected {
				t.Errorf("WordCount(%q) = %d; expected %d", tt.input, result, tt.expected)
			}
		})
	}
}

// BenchmarkReverse demonstrates benchmarking string operations
func BenchmarkReverse(b *testing.B) {
	input := "The quick brown fox jumps over the lazy dog"
	for i := 0; i < b.N; i++ {
		Reverse(input)
	}
}

// BenchmarkIsPalindrome demonstrates benchmarking complex operations
func BenchmarkIsPalindrome(b *testing.B) {
	input := "A man a plan a canal Panama"
	for i := 0; i < b.N; i++ {
		IsPalindrome(input)
	}
}

// ExampleReverse demonstrates example-based documentation
func ExampleReverse() {
	result := Reverse("hello")
	fmt.Println(result)
	// Output: olleh
}

// ExampleIsPalindrome demonstrates checking palindromes
func ExampleIsPalindrome() {
	result := IsPalindrome("racecar")
	fmt.Println(result)
	// Output: true
}
