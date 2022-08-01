package app

import (
	"encoding/csv"
	"flag"
	"fmt"
	"math"
	"os"
	"strings"
)

type config struct {
	fileName string
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

	result := runQuiz(problems)

	showScore(result)
}

func setup() config {
	defaults := config{
		fileName: "./data/problems.csv",
	}

	csvFileName := flag.String("csv", defaults.fileName, "a csv file in the format of 'question,answwer'")
	flag.Parse()

	return config{
		fileName: *csvFileName,
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

func runQuiz(problems []problem) score {
	s := score{
		points: 0,
		max:    len(problems),
	}

	for i, p := range problems {
		if askQuestion(i, p) {
			s.points++
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

func Exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
