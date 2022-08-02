package app

import (
	crypto_rand "crypto/rand"
	"encoding/binary"
	"fmt"
	math_rand "math/rand"
	"os"
)

type appInfo struct {
	name       string
	version    string
	gitVersion string
	gitHash    string
	buildTime  string
}

var metadata = appInfo{
	name:       "Gophercise Quiz App",
	version:    "0.0.0",
	gitVersion: "0.0.0",
	gitHash:    "abcdef",
	buildTime:  "1970-01-01T12:00:00Z",
}

type config struct {
	fileName  string
	timeLimit int
	shuffle   bool
}

var defaults = config{
	fileName:  "./data/problems.csv",
	timeLimit: 30,
	shuffle:   false,
}

// SetVersion is used my the main package to pass version information to the app package.
func SetVersion(slice ...string) {
	if slice[0] != "" {
		metadata.version = slice[0]
	}
	if slice[1] != "" {
		metadata.gitVersion = slice[1]
	}
	if slice[2] != "" {
		metadata.gitHash = slice[2]
	}
	if slice[3] != "" {
		metadata.buildTime = slice[3]
	}
}

func showVersion() {
	fmt.Println()
	fmt.Printf("%s Version: %s\n", metadata.name, metadata.version)
	fmt.Printf("git version: %s\n", metadata.gitVersion)
	fmt.Printf("   git hash: %s\n", metadata.gitHash)
	fmt.Printf(" build time: %s\n", metadata.buildTime)
	fmt.Println()
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

// Exit is used to prematurely end the application with an exit code and message to stdout.
func Exit(code int, msg string) {
	fmt.Println(msg)
	os.Exit(code)
}
