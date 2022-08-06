[![Go Report Card](https://goreportcard.com/badge/github.com/vpayno/gophercises-quizgame)](https://goreportcard.com/report/github.com/vpayno/gophercises-quizgame)
[![Go Workflow](https://github.com/vpayno/gophercises-quizgame/actions/workflows/go.yml/badge.svg)](https://github.com/vpayno/gophercises-quizgame/actions/workflows/go.yml)
[![Bash Workflow](https://github.com/vpayno/gophercises-quizgame/actions/workflows/bash.yml/badge.svg)](https://github.com/vpayno/gophercises-quizgame/actions/workflows/bash.yml)
[![Git Workflow](https://github.com/vpayno/gophercises-quizgame/actions/workflows/git.yml/badge.svg)](https://github.com/vpayno/gophercises-quizgame/actions/workflows/git.yml)
[![Link Check Workflow](https://github.com/vpayno/gophercises-quizgame/actions/workflows/links.yml/badge.svg)](https://github.com/vpayno/gophercises-quizgame/actions/workflows/links.yml)

![Coverage](./reports/.octocov-coverage.svg)
![Code2Test Ratio](./reports/.octocov-ratio.svg)

# Gophercises Quiz Game Implementation

## Gophercises Info

- [Website](https://courses.calhoun.io/lessons/les_goph_01)
- [GitHub](https://github.com/gophercises/quiz)

## How to Install

Using `go install`

```
$ go install github.com/vpayno/gophercises-quizgame/cmd/gophercises-quizgame@latest
```

or

```
$ git clone https://github.com/vpayno/gophercises-quizgame.git
$ cd gophercises-quizgame
$ make install
```

## Usage

```
$ go run ./cmd/gophercises-quizgame/gophercises-quizgame.go --help
Usage of /tmp/go-build172411189/b001/exe/gophercises-quizgame:
  -csv string
        a csv file in the format of 'question,answwer' (default "./data/problems.csv")
  -limit int
        the time limit for the quiz in seconds (default 30)
  -shuffle
        shuffle the questions
  -version
        show the app version
```

## How to Play

You have 30 seconds to answer all the questions.

```
$ go run ./cmd/gophercises-quizgame/gophercises-quizgame.go
Gophercise Quiz App Version 0.0.0


You have 30 seconds to answer 12 question.

1) 5+5 =
Time's up!

You scored 0 out of 12 points (0%).
```

```
$ go run ./cmd/gophercises-quizgame/gophercises-quizgame.go
Gophercise Quiz App Version 0.0.0


You have 30 seconds to answer 12 question.

1) 5+5 = 10
2) 1+1 = 0
3) 8+3 = 11
4) 1+2 = 3
5) 8+6 = 14
6) 3+1 = 4
7) 1+4 = 5
8) 5+1 = 6
9) 2+3 = 5
10) 3+3 = 6
11) 2+4 = 6
12) 5+2 = 7

You scored 11 out of 12 points (92%).
```
