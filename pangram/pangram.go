// Package pangram helps you with pangram stuff.
package pangram

import "strings"

const testVersion = 1

// IsPangram tells you if a string has all letters from a to z.
func IsPangram(s string) bool {
	var letters [26]bool

	s = strings.ToUpper(s)

	for _, w := range s {
		i := int(w) - 65
		if i >= 0 && i <= 25 {
			letters[i] = true
		}
	}

	for i := 0; i <= 25; i++ {
		if !letters[i] {
			return false
		}
	}

	return true
}
