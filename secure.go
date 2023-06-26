package rndpass

import (
	"crypto/rand"
	"fmt"
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

// MustSecureCode will generate a code similar to SecureCode, and will panic
// should the random generator fail to generate randomness (which is the proper
// action to take in this case).
func MustSecureCode(ln int, set string, rnd io.Reader) string {
	res, err := SecureCode(ln, set, rnd)
	if err != nil {
		panic(fmt.Sprintf("random generator failed: %s", err))
	}
	return res
}
