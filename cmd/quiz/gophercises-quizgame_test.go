package main

import (
	"fmt"
	"os"
	"testing"
)

// The functions in main() are already tested. Just running them together with zero test questions.
func TestMain(t *testing.T) {
	fileName := "../../test/data/problems-0of0.csv"
	timeLimit := 1

	os.Args = []string{"test", "-csv", fileName, "-limit", fmt.Sprintf("%d", timeLimit), "-shuffle"}

	main()
}
