// string_utils.go - String utilities for testing various scenarios
package main

import (
	"regexp"
	"strings"
	"unicode"
)

// StringUtils provides various string manipulation functions
type StringUtils struct{}

// NewStringUtils creates a new StringUtils instance
func NewStringUtils() *StringUtils {
	return &StringUtils{}
}

// Reverse reverses a string
func (s *StringUtils) Reverse(str string) string {
	runes := []rune(str)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// IsPalindrome checks if a string is a palindrome (case-insensitive)
func (s *StringUtils) IsPalindrome(str string) bool {
	cleaned := strings.ToLower(regexp.MustCompile(`[^a-zA-Z0-9]`).ReplaceAllString(str, ""))
	return cleaned == s.Reverse(cleaned)
}

// CountWords counts the number of words in a string
func (s *StringUtils) CountWords(str string) int {
	if str == "" {
		return 0
	}
	return len(strings.Fields(str))
}

// Capitalize capitalizes the first letter of each word
func (s *StringUtils) Capitalize(str string) string {
	return strings.Title(strings.ToLower(str))
}

// IsValidEmail performs basic email validation
func (s *StringUtils) IsValidEmail(email string) bool {
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return emailRegex.MatchString(email)
}

// ExtractNumbers extracts all numbers from a string
func (s *StringUtils) ExtractNumbers(str string) []string {
	numberRegex := regexp.MustCompile(`\d+`)
	return numberRegex.FindAllString(str, -1)
}

// RemoveSpaces removes all whitespace from a string
func (s *StringUtils) RemoveSpaces(str string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsSpace(r) {
			return -1
		}
		return r
	}, str)
}

// TruncateWithEllipsis truncates a string to a specified length and adds ellipsis
func (s *StringUtils) TruncateWithEllipsis(str string, maxLength int) string {
	if len(str) <= maxLength {
		return str
	}
	if maxLength <= 3 {
		return str[:maxLength]
	}
	return str[:maxLength-3] + "..."
}

// ContainsAny checks if string contains any of the provided substrings
func (s *StringUtils) ContainsAny(str string, substrings []string) bool {
	for _, substr := range substrings {
		if strings.Contains(str, substr) {
			return true
		}
	}
	return false
}

// GetInitials returns the initials from a full name
func (s *StringUtils) GetInitials(fullName string) string {
	words := strings.Fields(strings.TrimSpace(fullName))
	if len(words) == 0 {
		return ""
	}
	
	var initials strings.Builder
	for _, word := range words {
		if len(word) > 0 {
			initials.WriteString(strings.ToUpper(string(word[0])))
		}
	}
	return initials.String()
}
