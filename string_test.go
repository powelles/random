package random

import "testing"

func TestStringLength(t *testing.T) {
	for i := 1; i < 101; i++ {
		r := String(i)
		if len(r) != i {
			t.Error("String length not as specified")
		}
	}
}

func TestStringZero(t *testing.T) {
	r := String(0)
	if len(r) != 0 || r != "" {
		t.Error("Zero length string error")
	}
}

func TestStringRandomness(t *testing.T) {
	var s []string
	for i := 0; i < 1000; i++ {
		r := String(10)
		s = append(s, r)
	}

	for i1, v1 := range s {
		for i2, v2 := range s {
			if i1 != i2 && v1 == v2 {
				t.Error("Duplicate strings found")
			}
		}
	}
}

func TestStringChars(t *testing.T) {
	for i := 1; i < 1001; i++ {
		r := String(100)
		for _, v1 := range r {
			m := false
			for _, v2 := range chars {
				if v1 == v2 {
					m = true
				}
			}
			if !m {
				t.Error("Invalid charachter generated")
			}
		}
	}
}
