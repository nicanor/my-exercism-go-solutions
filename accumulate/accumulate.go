// Package accumulate has a nice function in it.
package accumulate

const testVersion = 1

// Accumulate is a map for strings.
func Accumulate(xs []string, f func(string) string) []string {
	ys := make([]string, len(xs))
	for i, v := range xs {
		ys[i] = f(v)
	}
	return ys
}
