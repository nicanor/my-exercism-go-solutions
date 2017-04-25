// Package scrabble calculates scrabble things.
package scrabble

import "strings"

const testVersion = 4

// Score calculates word scrabble points.
func Score(s string) int {
	s = strings.ToUpper(s)
	var sum int
	for _, r := range s {
		sum += runeScore(r)
	}
	return sum
}

func runeScore(r rune) int {
	switch {
	case r == 'A' || r == 'E' || r == 'I' || r == 'O' || r == 'U':
		return 1
	case r == 'L' || r == 'N' || r == 'R' || r == 'S' || r == 'T':
		return 1
	case r == 'D' || r == 'G':
		return 2
	case r == 'B' || r == 'C' || r == 'M' || r == 'P':
		return 3
	case r == 'F' || r == 'H' || r == 'V' || r == 'W' || r == 'Y':
		return 4
	case r == 'K':
		return 5
	case r == 'J' || r == 'X':
		return 8
	case r == 'Q' || r == 'Z':
		return 10
	default:
		return 0
	}

}
