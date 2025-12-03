package stringutil

import (
	"strings"
	"unicode"
)

// Reverse returns the reverse of a string
func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// IsPalindrome checks if a string is a palindrome (case-insensitive)
func IsPalindrome(s string) bool {
	s = strings.ToLower(s)
	// Remove spaces and punctuation
	cleaned := ""
	for _, r := range s {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			cleaned += string(r)
		}
	}
	return cleaned == Reverse(cleaned)
}

// CountVowels counts the number of vowels in a string
func CountVowels(s string) int {
	vowels := "aeiouAEIOU"
	count := 0
	for _, char := range s {
		if strings.ContainsRune(vowels, char) {
			count++
		}
	}
	return count
}

// ToTitle converts a string to title case
func ToTitle(s string) string {
	return strings.Title(s)
}

// WordCount returns the number of words in a string
func WordCount(s string) int {
	words := strings.Fields(s)
	return len(words)
}
