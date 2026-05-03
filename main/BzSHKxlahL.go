package main
import (
	"fmt"
	"strings"
)
func Reverse(s string) string {
	return reverseHelper(s, len(s)-1)
}
func reverseHelper(s string, index int) string {
	if index < 0 {
		return ""
	}
	return string(s[index]) + reverseHelper(s, index-1)
}
func ToUpper(s string) string {
	return strings.ToUpper(s)
}
func ToLower(s string) string {
	return strings.ToLower(s)
}
func IsPalindrome(s string) bool {
	rev := Reverse(s)
	return s == rev
}
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
func main() {
	example := "Hello, World!"
	fmt.Println("Original:", example)
	fmt.Println("Reversed:", Reverse(example))
	fmt.Println("Uppercase:", ToUpper(example))
	fmt.Println("Lowercase:", ToLower(example))
	fmt.Println("Is Palindrome:", IsPalindrome(example))
	fmt.Println("Vowel Count:", CountVowels(example))
}
