package types

import (
	"crypto/rand"
	"fmt"
)

type Hash [32]uint8 // 32 byte unsigned uint8

func (h Hash) IsZero() bool {
	for i := 0; i < 32; i++ {
		if h[i] != 0 {
			return false
		}
	}
	return true
}

// To hash the block(sha-256) and its going to return byte slice
func HashFromBytes(b []byte) Hash {
	if len(b) != 32 {
		msg := fmt.Sprintf("given bytes with length: %d should be 32", len(b))
		panic(msg) // system cannot continue
	}

	var value [32]uint8
	for i := 0; i < 32; i++ {
		value[i] = b[i]
	}

	return Hash(value)
}

func RandomBytes(size int) []byte {
	token := make([]byte, size)
	rand.Read(token)
	return token
}

func RandomHash() Hash {
	return HashFromBytes(RandomBytes(32))
}
