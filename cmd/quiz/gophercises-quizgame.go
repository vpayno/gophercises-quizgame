package main

import (
	app "github.com/vpayno/gophercises-quizgame/v2/internal/app/quiz"
)

var version string
var gitVersion string
var gitHash string
var buildTime string

func init() {
	app.InitRandSeed()
	app.SetVersion(version, gitVersion, gitHash, buildTime)
}

func main() {
	app.RunApp()
}
