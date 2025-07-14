// string_utils_test.go - Demonstrates various assertion types and test patterns
package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestStringUtilsReverse tests string reversal
func TestStringUtilsReverse(t *testing.T) {
	utils := NewStringUtils()

	testCases := []struct {
		name     string
		input    string
		expected string
	}{
		{"simple string", "hello", "olleh"},
		{"empty string", "", ""},
		{"single character", "a", "a"},
		{"palindrome", "racecar", "racecar"},
		{"with spaces", "hello world", "dlrow olleh"},
		{"with unicode", "café", "éfac"},
		{"numbers", "12345", "54321"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := utils.Reverse(tc.input)
			assert.Equal(t, tc.expected, result)
		})
	}
}

// TestStringUtilsIsPalindrome tests palindrome detection
func TestStringUtilsIsPalindrome(t *testing.T) {
	utils := NewStringUtils()

	// Test true cases
	palindromes := []string{
		"racecar",
		"A man a plan a canal Panama",
		"race a car",
		"",
		"a",
		"Madam",
		"Was it a car or a cat I saw?",
		"12321",
	}

	for _, p := range palindromes {
		t.Run("palindrome_"+p, func(t *testing.T) {
			assert.True(t, utils.IsPalindrome(p), "Should be palindrome: %s", p)
		})
	}

	// Test false cases
	notPalindromes := []string{
		"hello",
		"world",
		"test case",
		"not a palindrome",
		"12345",
	}

	for _, np := range notPalindromes {
		t.Run("not_palindrome_"+np, func(t *testing.T) {
			assert.False(t, utils.IsPalindrome(np), "Should not be palindrome: %s", np)
		})
	}
}

// TestStringUtilsCountWords tests word counting
func TestStringUtilsCountWords(t *testing.T) {
	utils := NewStringUtils()

	testCases := []struct {
		name     string
		input    string
		expected int
	}{
		{"empty string", "", 0},
		{"single word", "hello", 1},
		{"two words", "hello world", 2},
		{"multiple spaces", "hello    world", 2},
		{"leading/trailing spaces", "  hello world  ", 2},
		{"tabs and newlines", "hello\tworld\ntest", 3},
		{"punctuation", "hello, world!", 2},
		{"numbers", "123 456 789", 3},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := utils.CountWords(tc.input)
			assert.Equal(t, tc.expected, result, "Input: %q", tc.input)
		})
	}
}

// TestStringUtilsCapitalize tests capitalization
func TestStringUtilsCapitalize(t *testing.T) {
	utils := NewStringUtils()

	testCases := map[string]string{
		"hello world":           "Hello World",
		"HELLO WORLD":          "Hello World",
		"hELLo WoRLd":          "Hello World",
		"":                     "",
		"a":                    "A",
		"hello-world test_case": "Hello-World Test_Case",
		"123 test":             "123 Test",
	}

	for input, expected := range testCases {
		t.Run(input, func(t *testing.T) {
			result := utils.Capitalize(input)
			assert.Equal(t, expected, result)
		})
	}
}

// TestStringUtilsIsValidEmail tests email validation
func TestStringUtilsIsValidEmail(t *testing.T) {
	utils := NewStringUtils()

	// Valid emails
	validEmails := []string{
		"test@example.com",
		"user.name@domain.co.uk",
		"user+tag@example.org",
		"user123@test-domain.com",
		"a@b.co",
	}

	for _, email := range validEmails {
		t.Run("valid_"+email, func(t *testing.T) {
			assert.True(t, utils.IsValidEmail(email), "Should be valid: %s", email)
		})
	}

	// Invalid emails
	invalidEmails := []string{
		"",
		"invalid",
		"@example.com",
		"test@",
		"test@.com",
		"test.example.com",
		"test@example",
		"test @example.com",
		"test@exa mple.com",
	}

	for _, email := range invalidEmails {
		t.Run("invalid_"+email, func(t *testing.T) {
			assert.False(t, utils.IsValidEmail(email), "Should be invalid: %s", email)
		})
	}
}

// TestStringUtilsExtractNumbers tests number extraction
func TestStringUtilsExtractNumbers(t *testing.T) {
	utils := NewStringUtils()

	testCases := []struct {
		name     string
		input    string
		expected []string
	}{
		{"no numbers", "hello world", []string{}},
		{"single number", "hello123world", []string{"123"}},
		{"multiple numbers", "test123and456", []string{"123", "456"}},
		{"numbers with spaces", "test 123 and 456", []string{"123", "456"}},
		{"only numbers", "123456", []string{"123456"}},
		{"mixed", "abc123def456ghi", []string{"123", "456"}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := utils.ExtractNumbers(tc.input)
			assert.Equal(t, tc.expected, result)
		})
	}
}

// TestStringUtilsRemoveSpaces tests space removal
func TestStringUtilsRemoveSpaces(t *testing.T) {
	utils := NewStringUtils()

	testCases := map[string]string{
		"hello world":     "helloworld",
		"  test  ":        "test",
		"a b c d":         "abcd",
		"no-spaces":       "no-spaces",
		"":                "",
		"   ":             "",
		"hello\tworld\n":  "helloworld",
	}

	for input, expected := range testCases {
		t.Run(input, func(t *testing.T) {
			result := utils.RemoveSpaces(input)
			assert.Equal(t, expected, result)
		})
	}
}

// TestStringUtilsTruncateWithEllipsis tests truncation
func TestStringUtilsTruncateWithEllipsis(t *testing.T) {
	utils := NewStringUtils()

	testCases := []struct {
		name      string
		input     string
		maxLength int
		expected  string
	}{
		{"shorter than max", "hello", 10, "hello"},
		{"equal to max", "hello", 5, "hello"},
		{"longer than max", "hello world", 8, "hello..."},
		{"very short max", "hello", 3, "hel"},
		{"empty string", "", 5, ""},
		{"max length 0", "hello", 0, ""},
		{"single character", "a", 1, "a"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := utils.TruncateWithEllipsis(tc.input, tc.maxLength)
			assert.Equal(t, tc.expected, result)
			assert.LessOrEqual(t, len(result), tc.maxLength, "Result should not exceed max length")
		})
	}
}

// TestStringUtilsContainsAny tests substring checking
func TestStringUtilsContainsAny(t *testing.T) {
	utils := NewStringUtils()

	t.Run("contains one", func(t *testing.T) {
		result := utils.ContainsAny("hello world", []string{"world", "test"})
		assert.True(t, result)
	})

	t.Run("contains multiple", func(t *testing.T) {
		result := utils.ContainsAny("hello world test", []string{"world", "test"})
		assert.True(t, result)
	})

	t.Run("contains none", func(t *testing.T) {
		result := utils.ContainsAny("hello world", []string{"foo", "bar"})
		assert.False(t, result)
	})

	t.Run("empty substrings", func(t *testing.T) {
		result := utils.ContainsAny("hello world", []string{})
		assert.False(t, result)
	})

	t.Run("empty string", func(t *testing.T) {
		result := utils.ContainsAny("", []string{"test"})
		assert.False(t, result)
	})
}

// TestStringUtilsGetInitials tests initial extraction
func TestStringUtilsGetInitials(t *testing.T) {
	utils := NewStringUtils()

	testCases := map[string]string{
		"John Doe":           "JD",
		"John Michael Doe":   "JMD",
		"john doe":           "JD",
		"  John   Doe  ":     "JD",
		"":                   "",
		"   ":                "",
		"Madonna":            "M",
		"Jean-Claude Van Damme": "JCVD",
		"A B C D E":          "ABCDE",
	}

	for input, expected := range testCases {
		t.Run(input, func(t *testing.T) {
			result := utils.GetInitials(input)
			assert.Equal(t, expected, result)
		})
	}
}

// TestStringUtilsEdgeCases tests various edge cases
func TestStringUtilsEdgeCases(t *testing.T) {
	utils := NewStringUtils()

	t.Run("unicode handling", func(t *testing.T) {
		// Test with emoji
		result := utils.Reverse("hello 👋 world")
		assert.Contains(t, result, "👋")

		// Test initials with unicode
		initials := utils.GetInitials("José María Ñuñez")
		assert.Equal(t, "JMÑ", initials)
	})

	t.Run("very long strings", func(t *testing.T) {
		longString := string(make([]byte, 10000))
		for i := range longString {
			longString = longString[:i] + "a" + longString[i+1:]
		}
		
		result := utils.Reverse(longString)
		assert.Len(t, result, 10000)
		assert.Equal(t, "a", string(result[0]))
	})

	t.Run("nil safety", func(t *testing.T) {
		// These shouldn't panic
		assert.NotPanics(t, func() {
			utils.Reverse("")
			utils.CountWords("")
			utils.GetInitials("")
			utils.ContainsAny("", nil)
		})
	})
}

// BenchmarkStringUtilsReverse benchmarks string reversal
func BenchmarkStringUtilsReverse(b *testing.B) {
	utils := NewStringUtils()
	testString := "The quick brown fox jumps over the lazy dog"
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		utils.Reverse(testString)
	}
}

// BenchmarkStringUtilsIsPalindrome benchmarks palindrome checking
func BenchmarkStringUtilsIsPalindrome(b *testing.B) {
	utils := NewStringUtils()
	testString := "A man a plan a canal Panama"
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		utils.IsPalindrome(testString)
	}
}

// TestStringUtilsWithRequire demonstrates using require for critical assertions
func TestStringUtilsWithRequire(t *testing.T) {
	utils := NewStringUtils()
	
	// Use require for critical setup that other tests depend on
	require.NotNil(t, utils, "StringUtils instance should not be nil")
	
	// If this fails, the test stops here
	result := utils.Reverse("test")
	require.NotEmpty(t, result, "Reverse should return non-empty result")
	
	// This assertion runs only if the above require passes
	assert.Equal(t, "tset", result)
}
