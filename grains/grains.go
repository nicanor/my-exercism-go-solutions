// Package grains solves grains chess problem.
package grains

import "errors"

const testVersion = 1

// Returns total number of grains.
func Total() uint64 {
	return 18446744073709551615
}

// Returns numbers of grains for given square.
func Square(i int) (uint64, error) {
	if i <= 0 || i > 64 {
		return 0, errors.New("Invalid input")
	}
	return uint64(1) << uint64(i-1), nil
}
