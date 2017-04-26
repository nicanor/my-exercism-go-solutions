// Package pythagorean is useful for Pythagorean triplets calculations.
package pythagorean

type Triplet [3]int

const testVersion = 1

// Range returns a list of all Pythagorean triplets with sides in the
// range min to max inclusive.
func Range(min, max int) []Triplet {
	var t []Triplet
	for i := min; i < max-1; i++ {
		var square int = i * i
		for j := i + 1; j < max; j++ {
			var squarej = j * j
			for k := j + 1; k <= max; k++ {
				if k*k == square+squarej {
					t = append(t, Triplet{i, j, k})
				}
			}
		}
	}

	return t
}

// Sum returns a list of all Pythagorean triplets where the sum a+b+c
// (the perimeter) is equal to p.
func Sum(p int) []Triplet {
	var t []Triplet = Range(1, p)
	var w []Triplet
	for i := 1; i < len(t); i++ {
		var x = t[i]
		if x[0]+x[1]+x[2] == p {
			w = append(w, x)
		}
	}
	return w
}
