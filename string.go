package random

import "unicode"

// String takes an int as string length and returns a random alphanumeric string of the requested length.
func String(l int) string {
	b := make([]byte, l)
	for i := 0; i < l; i++ {
		r := newRune()
		for !unicode.IsNumber(r) && !unicode.IsLetter(r) {
			r = newRune()
		}
		b[i] = byte(r)
	}
	return string(b)
}
