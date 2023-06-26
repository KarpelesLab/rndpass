package rndpass

import (
	"math/rand"
	"sync"
	"time"
)

const (
	RangeFull       = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	RangePassword   = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789&~#([-|_^@)]{}$*,?;.:/!<>"
	RangeAlnumLower = "abcdefghijklmnopqrstuvwxyz0123456789"
	RangeNumeric    = "0123456789"
)

// we might want to switch this to crypto/rand
var (
	codeRnd     *rand.Rand
	codeRndOnce sync.Once
)

// Code generates a random string made of ln characters taken from set. For
// example to generate a random password of 32 characters, you can simply
// call Code(32, rndpass.RangePassword). This method uses golang's math/rand
// random generator and is not considered secure, but can come in handy in
// some cases as it cannot fail.
//
// If you need higher quality randomness, please use SecureCode()
func Code(ln int, set string) string {
	codeRndOnce.Do(func() {
		// perform random number generator only at first run
		codeRnd = rand.New(rand.NewSource(time.Now().UnixNano()))
	})

	v := make([]byte, 0, ln)
	x := len(set)

	for ln > 0 {
		r := codeRnd.Intn(x)
		v = append(v, set[r])
		ln -= 1
	}
	return string(v)
}
