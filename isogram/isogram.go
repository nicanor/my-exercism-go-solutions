
package isogram

import "strings"

const testVersion = 1

// IsIsogram determines if a word or phrase is an isogram.
func IsIsogram(s string) bool {
	m := map[rune]bool{}
	s = strings.ToLower(s)
	for _, r := range s {
		if r != '-' && r != ' ' {
			if m[r] {
				return false
			}
			m[r] = true
		}
	}
	return true
}
