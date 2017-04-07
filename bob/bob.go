// Package bob is useful for bob.
package bob

import (
	"strings"
	"unicode"
)

const testVersion = 2

// Hey tooks a string and returns bob's answer.
func Hey(s string) string {
	s = strings.TrimSpace(s)
	if s == "" {
		return "Fine. Be that way!"
	}
	if allUpper(s) {
		return "Whoa, chill out!"
	}
	if strings.HasSuffix(s, "?") {
		return "Sure."
	}
	return "Whatever."
}

func allUpper(s string) bool {
	var anyUpper bool
	for _, r := range s {
		if unicode.IsUpper(r) {
			anyUpper = true
		}
		if unicode.IsLower(r) {
			return false
		}
	}
	return anyUpper
}
