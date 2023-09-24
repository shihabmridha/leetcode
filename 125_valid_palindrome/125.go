package validpalindrome

import (
	"regexp"
	"strings"
)

func isPalindrome(s string) bool {
	s = regexp.MustCompile(`[^a-zA-Z0-9]+`).ReplaceAllString(s, "")
	s = strings.ToLower(s)

	ln := len(s)

	isPalindrome := true

	for i, j := 0, ln-1; i <= j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			isPalindrome = false
			break
		}
	}

	return isPalindrome
}
