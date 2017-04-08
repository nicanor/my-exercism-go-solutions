// Package secret is not documented.
package secret

const testVersion = 1

var normalMasks = []uint{1, 2, 4, 8}
var invertMasks = []uint{8, 4, 2, 1}

// Handshake is not documented.
func Handshake(i uint) []string {
	if i&16 == 0 {
		return handshake(i, normalMasks)
	}
	return handshake(i, invertMasks)
}

func handshake(i uint, masks []uint) []string {
	var xs []string
	for _, mask := range masks {
		if mask&i == mask {
			xs = append(xs, shake(mask))
		}
	}
	return xs
}

func shake(i uint) string {
	switch i {
	case 1:
		return "wink"
	case 2:
		return "double blink"
	case 4:
		return "close your eyes"
	case 8:
		return "jump"
	}
	return ""
}
