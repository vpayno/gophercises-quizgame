[![Go Workflow](https://github.com/vpayno/gophercises-quizgame/actions/workflows/go.yml/badge.svg)](https://github.com/vpayno/gophercises-quizgame/actions/workflows/go.yml)
[![Bash Workflow](https://github.com/vpayno/gophercises-quizgame/actions/workflows/bash.yml/badge.svg)](https://github.com/vpayno/gophercises-quizgame/actions/workflows/bash.yml)
[![Git Workflow](https://github.com/vpayno/gophercises-quizgame/actions/workflows/git.yml/badge.svg)](https://github.com/vpayno/gophercises-quizgame/actions/workflows/git.yml)

# Gophercises Quiz Game Implementation

## Gophercises Info

- [Website](https://courses.calhoun.io/lessons/les_goph_01)
- [GitHub](https://github.com/gophercises/quiz)

## How to Install

For `tag`, use `main`, `latest` or a tagged version.

```
$ go install github.com/vpayno/gophercises-quizgame/v2/cmd/quiz@tag
```

## Usage

```
$ go run ./cmd/quiz/gophercises-quizgame.go --help
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
$ go run ./cmd/quiz/gophercises-quizgame.go
Gophercise Quiz App Version 0.0.0


You have 30 seconds to answer 12 question.

1) 5+5 =
Time's up!

You scored 0 out of 12 points (0%).
```

```
$ go run ./cmd/quiz/gophercises-quizgame.go
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
