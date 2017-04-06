// Package hamming has functions for DNA calculations.
package hamming

import "errors"

const testVersion = 5

// Distance calculates Hamming difference between two DNA strands.
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return -1, errors.New("error")
	}
	count := 0

	for i, len := 0, len(a); i < len; i++ {
		if a[i] != b[i] {
			count += 1
		}
	}

	return count, nil
}
