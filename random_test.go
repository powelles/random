package random

import "testing"

func TestNewInt(t *testing.T) {
	for i := 0; i < 10000; i++ {
		n := newInt(100)
		if n < 0 || n > 100 {
			t.Error("New Int Out of Range")
		}
	}
}

func TestNewRuneRandomness(t *testing.T) {
	for i := 0; i < 100; i++ {
		a, b, c := newRune(), newRune(), newRune()
		if a == b && b == c {
			t.Error("Duplicate runes")
		}
	}
}
