package main

import (
	"fmt"
	"os"
	"testing"

	app "github.com/vpayno/gophercises-quizgame/internal/app/quiz"
)

// The functions in main() are already tested. Just running them together with zero test questions.
func TestMain(t *testing.T) {
	fileName := "../../test/data/problems-0of0.csv"
	timeLimit := 1

	os.Args = []string{"test", "-csv", fileName, "-limit", fmt.Sprintf("%d", timeLimit), "-shuffle"}

	app.SetVersion(version)
	main()
}
