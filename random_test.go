package random

import "testing"

func TestNewIntRange(t *testing.T) {
	l := 100
	for i := 0; i < 10000; i++ {
		n := newInt(l)
		if n < 0 || n > l {
			t.Error("New Int Out of Range")
		}
	}
}

func TestNewIntValues(t *testing.T) {
	l := 100
	for i := 0; i <= l; i++ {
		m := false
		for ii := 0; ii < l*100; ii++ {
			n := newInt(l)
			if i == n {
				m = true
				break
			}
		}
		if !m {
			t.Error("Missing New Int Value")
		}
	}

}

func TestNewRuneValues(t *testing.T) {
	for _, v := range chars {
		m := false
		for i := 0; i < 1000; i++ {
			r := newRune()
			if r == v {
				m = true
				break
			}
		}

		if !m {
			t.Error("Missing rune in char range")
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
