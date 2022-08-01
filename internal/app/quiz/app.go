package app

import (
	"encoding/csv"
	"flag"
	"fmt"
	"math"
	"os"
	"strings"
	"time"
)

type config struct {
	fileName  string
	timeLimit int
}

type problem struct {
	question string
	answer   string
}

type score struct {
	points int
	max    int
}

func (this *score) rate() int {
	return int(math.Round(float64(this.points) / float64(this.max) * 100))
}

type quizData [][]string

func RunApp() {
	fmt.Println("quiz app")
	fmt.Println()

	c := setup()

	data := loadData(c)

	problems := parseLines(data)

	timer := createTimer(c)

	result := runQuiz(c, problems, timer)

	showScore(result)
}

func setup() config {
	defaults := config{
		fileName:  "./data/problems.csv",
		timeLimit: 30,
	}

	csvFileName := flag.String("csv", defaults.fileName, "a csv file in the format of 'question,answwer'")
	timeLimit := flag.Int("limit", defaults.timeLimit, "the time limit for the quiz in seconds")
	flag.Parse()

	return config{
		fileName:  *csvFileName,
		timeLimit: *timeLimit,
	}
}

func loadData(c config) quizData {
	file, err := os.Open(c.fileName)
	if err != nil {
		Exit(fmt.Sprintf("Failed to open the CSV file: %q\n", c.fileName))
	}

	defer file.Close()

	r := csv.NewReader(file)

	lines, err := r.ReadAll()
	if err != nil {
		Exit(fmt.Sprintf("Failed to parse the provided CSV file: %q\n", c.fileName))
	}

	return lines
}

func parseLines(lines quizData) []problem {
	result := make([]problem, len(lines))

	for i, line := range lines {
		result[i] = problem{
			question: strings.TrimSpace(line[0]),
			answer:   strings.TrimSpace(line[1]),
		}
	}

	return result
}

func runQuiz(c config, problems []problem, timer *time.Timer) score {
	s := score{
		points: 0,
		max:    len(problems),
	}

	fmt.Println()
	fmt.Printf("You have %d seconds to answer %d question.\n", c.timeLimit, s.max)
	fmt.Println()

	for i, p := range problems {
		answerCh := make(chan bool)
		go func() {
			answerCh <- askQuestion(i, p)
		}()

		select {
		case <-timer.C:
			fmt.Println()
			fmt.Println()
			fmt.Println("Time's up!")
			return s
		case response := <-answerCh:
			if response {
				s.points++
			}
		}
	}

	return s
}

func askQuestion(i int, p problem) bool {
	fmt.Printf("%d) %s = ", i+1, p.question)

	var response string
	fmt.Scanf("%s", &response)

	return response == p.answer
}

func showScore(s score) {
	fmt.Println()
	fmt.Printf("You scored %d out of %d points (%v%%).\n", s.points, s.max, s.rate())
	fmt.Println()
}

func createTimer(c config) *time.Timer {
	timer := time.NewTimer(time.Duration(c.timeLimit) * time.Second)
	return timer
}

func Exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
