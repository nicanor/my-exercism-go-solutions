// Package summultiples has functions for calculating the sum of all the multiples of particular numbers up to but not including a given number.

package summultiples

const testVersion = 1

// SumMultiples calculates the sum of all the multiples of particular numbers up to but not including a given number.
func SumMultiples(to int, divisors ...int) int {
	var total int
	for i := 1; i < to; i++ {
		if multiple(i, divisors) {
			total += i
		}
	}
	return total
}

func multiple(i int, divisors []int) bool {
	for _, divisor := range divisors {
		if i%divisor == 0 {
			return true
		}
	}
	return false
}
