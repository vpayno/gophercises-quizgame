package app

import (
	"encoding/csv"
	"flag"
	"fmt"
	"math"
	math_rand "math/rand"
	"os"
	"strings"
	"time"
)

type problem struct {
	question string
	answer   string
}

type score struct {
	points int
	max    int
}

func (s *score) rate() int {
	return int(math.Round(float64(s.points) / float64(s.max) * 100))
}

type quizData [][]string

func showBanner() {
	fmt.Println(metadata.name + " Version " + metadata.version)
	fmt.Println()
}

// RunApp is called my the main function. It's basically the main function of the app.
func RunApp() {
	c := setup()

	showBanner()

	data := loadData(c)

	problems := parseLines(data)

	timer := createTimer(c)

	result := runQuiz(c, problems, timer)

	showScore(result)
}

func setup() config {
	csvFileName := flag.String("csv", defaults.fileName, "a csv file in the format of 'question,answwer'")
	timeLimit := flag.Int("limit", defaults.timeLimit, "the time limit for the quiz in seconds")
	shuffle := flag.Bool("shuffle", defaults.shuffle, "shuffle the questions")
	version := flag.Bool("version", false, "show the app version")
	flag.Parse()

	if *version {
		showVersion()
		Exit(0, "")
	}

	return config{
		fileName:  *csvFileName,
		timeLimit: *timeLimit,
		shuffle:   *shuffle,
	}
}

func shuffleData(data quizData) {
	math_rand.Shuffle(len(data), func(i, j int) {
		data[i], data[j] = data[j], data[i]
	})
}

func loadData(c config) quizData {
	file, err := os.Open(c.fileName)
	if err != nil {
		Exit(1, fmt.Sprintf("Failed to open the CSV file: %q\n", c.fileName))
	}

	defer func() {
		if err := file.Close(); err != nil {
			fmt.Printf("Error closing file: %s\n", err)
		}
	}()

	r := csv.NewReader(file)

	lines, err := r.ReadAll()
	if err != nil {
		Exit(1, fmt.Sprintf("Failed to parse the provided CSV file: %q\n", c.fileName))
	}

	if c.shuffle {
		shuffleData(lines)
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
	for {
		fmt.Printf("%d) %s = ", i+1, p.question)

		var response string
		count, err := fmt.Scanf("%s", &response)

		if count == 1 && err == nil {
			return response == p.answer
		}

		fmt.Println("error enountered, restarting question")
	}
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
