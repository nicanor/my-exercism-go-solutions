// Package etl has functions for transforming scrabble scores.
package etl

import "strings"

const testVersion = 1

// Transform passes old scrable scores notation to new notation.
func Transform(table map[int][]string) map[string]int {
	m := make(map[string]int)
	for key, values := range table {
		for _, value := range values {
			m[strings.ToLower(value)] = key
		}
	}
	return m
}
