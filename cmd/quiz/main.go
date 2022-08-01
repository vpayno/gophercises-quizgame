package main

import (
	app "github.com/vpayno/gophercises-quizgame/v2/internal/app/quiz"
)

func init() {
	app.InitRandSeed()
}

func main() {
	app.RunApp()
}
