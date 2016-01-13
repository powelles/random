package random

import "testing"

func TestIntRandomness(t *testing.T) {
	for i := 0; i < 100; i++ {
		a, b, c := Int(100), Int(100), Int(100)
		if a == b && b == c {
			t.Error("Duplicate Ints")
		}
	}
}
