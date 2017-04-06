// Package for doing the raindrop thing.
package raindrops

import "fmt"

const testVersion = 2

// Convert does the raindrop thing.
func Convert(i int) string {
	x := ""

	if i%3 == 0 {
		x += "Pling"
	}

	if i%5 == 0 {
		x += "Plang"
	}

	if i%7 == 0 {
		x += "Plong"
	}

	if x == "" {
		x = fmt.Sprintf("%d", i)
	}

	return x
}
