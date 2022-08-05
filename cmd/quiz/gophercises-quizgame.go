package main

import (
	app "github.com/vpayno/gophercises-quizgame/internal/app/quiz"

	_ "embed"
)

//go:generate bash ../../scripts/go-generate-helper-git-version-info
//go:embed .version.txt
var version []byte

func init() {
	app.InitRandSeed(-1, false)
	app.SetVersion(version)
}

func main() {
	app.RunApp()
}
