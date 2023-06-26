package rndpass

import (
	"crypto/rand"
	"io"
	"math/big"
)

// SecureCode method works similarly to Code() except it uses a secure random
// source (if rnd is nil, crypto/rand.Reader is used by default) and may return
// an error if reading from the random source fails (it shouldn't). This method
// will be slower than Code() but comes with a much higher randomness quality.
func SecureCode(ln int, set string, rnd io.Reader) (string, error) {
	if rnd == nil {
		rnd = rand.Reader
	}

	v := make([]byte, 0, ln)
	x := big.NewInt(int64(len(set)))

	for ln > 0 {
		r, err := rand.Int(rnd, x)
		if err != nil {
			return "", err
		}
		v = append(v, set[r.Int64()])
		ln -= 1
	}
	return string(v), nil
}
