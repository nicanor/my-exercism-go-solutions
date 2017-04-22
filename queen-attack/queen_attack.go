// Package queenattack solves chess problems.
package queenattack

import "errors"

const testVersion = 2

// CanQueenAttack responds the question "are this queens in their respective attacking range?".
func CanQueenAttack(w, b string) (bool, error) {
	if w != b && isValid(w) && isValid(b) {
		return areAttacking(w, b), nil
	}
	return false, errors.New("Invalid positions")
}

func isValid(c string) bool {
	return len(c) == 2 &&
		c[0] >= 'a' && c[0] <= 'h' &&
		c[1] >= '1' && c[1] <= '8'
}

func areAttacking(w, b string) bool {
	return w[0] == b[0] || w[1] == b[1] ||
		w[0]-b[0] == w[1]-b[1] ||
		w[0]-b[0] == b[1]-w[1]
}
