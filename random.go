// Package random easily generates crypto random strings and ints.
package random

import (
	"crypto/rand"
	"math/big"
)

const chars string = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

func newInt(i int) int {
	r := big.NewInt(0)
	r, err := rand.Int(rand.Reader, big.NewInt(int64(i+1)))
	if err != nil {
		panic(err)
	}
	return int(r.Int64())
}

func newRune() rune {
	return rune(newInt(130))
}
