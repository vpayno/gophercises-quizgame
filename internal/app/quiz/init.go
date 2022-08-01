package app

import (
	crypto_rand "crypto/rand"
	"encoding/binary"
	math_rand "math/rand"
)

type config struct {
	fileName  string
	timeLimit int
	shuffle   bool
}

var defaults config = config{
	fileName:  "./data/problems.csv",
	timeLimit: 30,
	shuffle:   false,
}

// InitRandSeed seeds the random number library.
// This is better than just calling: `rand.Seed(time.Now().UnixNano())`
func InitRandSeed() {
	var b [8]byte

	_, err := crypto_rand.Read(b[:])
	if err != nil {
		panic("cannot seed math/rand package with cryptographically secure random number generator")
	}

	math_rand.Seed(int64(binary.LittleEndian.Uint64(b[:])))
}
