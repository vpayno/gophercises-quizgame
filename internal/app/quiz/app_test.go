package app

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
  Stdout testing code borrowed from Jon Calhoun's FizzBuzz example.
  https://courses.calhoun.io/lessons/les_algo_m01_08
  https://github.com/joncalhoun/algorithmswithgo.com/blob/master/module01/fizz_buzz_test.go
*/

func TestShowBanner(t *testing.T) {
	testStdout, writer, err := os.Pipe()
	if err != nil {
		t.Errorf("os.Pipe() err %v; want %v", err, nil)
	}

	osStdout := os.Stdout // keep backup of the real stdout
	os.Stdout = writer

	defer func() {
		// Undo what we changed when this test is done.
		os.Stdout = osStdout
	}()

	// It's a silly test but I need the practice.
	want := metadata.name + " Version " + metadata.version + "\n\n"

	// Run the function who's output we want to capture.
	showBanner()

	// Stop capturing stdout.
	writer.Close()

	var buf bytes.Buffer
	_, err = io.Copy(&buf, testStdout)
	if err != nil {
		t.Error(fmt.Sprint(err))
	}
	got := buf.String()
	if got != want {
		t.Errorf("showBanner(); want %q, got %q", want, got)
	}
}

func TestShowScore(t *testing.T) {
	testStdout, writer, err := os.Pipe()
	if err != nil {
		t.Errorf("os.Pipe() err %v; want %v", err, nil)
	}

	osStdout := os.Stdout // keep backup of the real stdout
	os.Stdout = writer

	defer func() {
		// Undo what we changed when this test is done.
		os.Stdout = osStdout
	}()

	// It's a silly test but I need the practice.
	s := score{
		points: 8,
		max:    10,
	}
	want := fmt.Sprintf("\nYou scored %d out of %d points (%v%%).\n\n", s.points, s.max, s.rate())

	// Run the function who's output we want to capture.
	showScore(s)

	// Stop capturing stdout.
	writer.Close()

	var buf bytes.Buffer
	_, err = io.Copy(&buf, testStdout)
	if err != nil {
		t.Error(fmt.Sprint(err))
	}
	got := buf.String()
	if got != want {
		t.Errorf("showBanner(); want %q, got %q", want, got)
	}
}

func TestCreateTimer(t *testing.T) {
	c := config{
		timeLimit: 1,
	}

	timer1 := createTimer(c)
	timer2 := createTimer(c)
	var got bool

	select {
	case <-timer1.C:
		got = true
	case <-timer2.C:
		got = true
	}

	if !got {
		t.Error("createTimer() didn't work")
	}
}

func areSlicesEqual(a, b quizData) bool {
	if len(a) != len(b) {
		return false
	}

	for i, v1 := range a {
		for j, v2 := range v1 {
			if v2 != b[i][j] {
				return false
			}
		}
	}

	return true
}

func TestShuffleData(t *testing.T) {
	slice := quizData{
		[]string{"one", "two"},
		[]string{"three", "four"},
		[]string{"five", "six"},
		[]string{"seven", "eight"},
		[]string{"nine", "ten"},
	}

	got := quizData{
		[]string{"one", "two"},
		[]string{"three", "four"},
		[]string{"five", "six"},
		[]string{"seven", "eight"},
		[]string{"nine", "ten"},
	}

	InitRandSeed(-1, false)
	shuffleData(got)

	if areSlicesEqual(slice, got) {
		t.Errorf("slice %#v is equal to slice %#v so it wasn't shuffled", slice, got)
	}
}

func TestLoadData(t *testing.T) {
	c := config{
		fileName: "../../../test/data/problems-12of12.csv",
	}

	data := loadData(c)

	if len(data) != 12 {
		t.Errorf("failed to read 12 lines from the data file")

		path, err := os.Getwd()
		if err != nil {
			log.Println(err)
		}
		t.Logf("cwd: %q\n", path)
	}
}

func TestLoadDataMissingFileAndInvalidData(t *testing.T) {
	c := config{
		fileName: "../../../test/data/does_not_exist.csv",
		shuffle:  true,
	}

	want := "Failed to open the CSV file: \"" + c.fileName + "\"\n\n"
	want += "Failed to parse the provided CSV file: \"" + c.fileName + "\"\n\n"
	want += "Error closing file: invalid argument\n"

	testStdout, writer, err := os.Pipe()
	if err != nil {
		t.Errorf("os.Pipe() err %v; want %v", err, nil)
	}

	osStdout := os.Stdout // keep backup of the real stdout
	os.Stdout = writer

	defer func() {
		// Undo what we changed when this test is done.
		os.Stdout = osStdout
	}()

	OSExitBackup := OSExit
	OSExit = func(code int) { _ = code }
	// It's not going to exit, it will return a value we don't want.
	data := loadData(c)
	OSExit = OSExitBackup

	// Stop capturing stdout.
	writer.Close()

	var buf bytes.Buffer
	_, err = io.Copy(&buf, testStdout)
	if err != nil {
		t.Error(fmt.Sprint(err))
	}
	got := buf.String()
	if got != want {
		t.Errorf("Exit(); want %q, got %q", want, got)
	}
	assert.Equal(t, len(data), 0)
}

func TestParseLines(t *testing.T) {
	c := config{
		fileName: "../../../test/data/problems-12of12.csv",
	}

	data := loadData(c)

	problems := parseLines(data)

	if len(data) != len(problems) {
		t.Errorf("loaded data (%d) and processed data (%d) don't have equal lengths", len(data), len(problems))

		path, err := os.Getwd()
		if err != nil {
			log.Println(err)
		}
		t.Logf("cwd: %q\n", path)
	}

	var s interface{} = problems
	if _, ok := s.([]problem); !ok {
		t.Errorf("process data, %T, isn't of type []app.problem", problems)
	}

	if problems[0].question != "5+5" && problems[0].answer != "10" {
		t.Errorf("problem slice doesn't have expected data: %#v", problems)
	}
}

func TestSetupFlags(t *testing.T) {
	want := config{
		fileName:  "test.csv",
		timeLimit: 1,
		shuffle:   true,
	}

	// -csv=test.csv -limit=1 -shuffle
	os.Args = []string{"test", "-csv", want.fileName, "-limit", fmt.Sprintf("%d", want.timeLimit), "-shuffle"}

	got := setup()

	if want.fileName != got.fileName {
		t.Errorf("setup() returned the wrong file name. want: %q, got %q", want.fileName, got.fileName)
	}

	if want.timeLimit != got.timeLimit {
		t.Errorf("setup() returned the wrong time limit. want: %d, got %d", want.timeLimit, got.timeLimit)
	}

	if want.shuffle != got.shuffle {
		t.Errorf("setup() returned the wrong shuffle value. want: %v, got %v", want.shuffle, got.shuffle)
	}

	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError) // flags are now reset
}

func TestSetupFlagVersion(t *testing.T) {
	testStdout, writer, err := os.Pipe()
	if err != nil {
		t.Errorf("os.Pipe() err %v; want %v", err, nil)
	}

	osStdout := os.Stdout // keep backup of the real stdout
	os.Stdout = writer

	defer func() {
		// Undo what we changed when this test is done.
		os.Stdout = osStdout
	}()

	want := "\n"
	want += fmt.Sprintf("%s Version: %s\n", metadata.name, metadata.version)
	want += fmt.Sprintf("git version: %s\n", metadata.gitVersion)
	want += fmt.Sprintf("   git hash: %s\n", metadata.gitHash)
	want += fmt.Sprintf(" build time: %s\n", metadata.buildTime)
	want += "\n"
	want += "\n"

	// -version
	os.Args = []string{"test", "-version"}
	OSExitBackup := OSExit
	OSExit = func(code int) { _ = code }
	// It's not going to exit, it will return a value we don't want.
	_ = setup()
	OSExit = OSExitBackup

	// Stop capturing stdout.
	writer.Close()

	var buf bytes.Buffer
	_, err = io.Copy(&buf, testStdout)
	if err != nil {
		t.Error(fmt.Sprint(err))
	}
	got := buf.String()
	if got != want {
		t.Errorf("Exit(); want %q, got %q", want, got)
	}

	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError) // flags are now reset
}

func TestAskQuestion(t *testing.T) {
	var stdin bytes.Buffer

	p := problem{
		question: "100+23",
		answer:   "123",
	}

	want := true

	bytes := []byte(p.answer)
	bytes = append(bytes, '\n')
	size, err := stdin.Write(bytes)

	got := askQuestion(0, p, &stdin)

	assert.NoError(t, err) // Yes, stdin.Write() always returns nil for error. It panics when it encounters an error.
	assert.Equal(t, len(p.answer+"\n"), size)
	assert.Equal(t, want, got)
}

func TestRunQuiz(t *testing.T) {
	var stdin bytes.Buffer

	c := config{
		fileName:  "../../../test/data/problems-1of1.csv",
		timeLimit: 1,
	}

	// tested in TestLoadData, assuming it's good
	data := loadData(c)

	// tested in TestParseLines, assuming it's good
	problems := parseLines(data)

	// tested in TestTimer, assuming it's good
	timer := createTimer(c)

	// Create our input stream.
	var answers string
	for _, p := range problems {
		answers += p.answer + "\n"
	}
	bytes := []byte(answers)
	size, err := stdin.Write(bytes)
	assert.NoError(t, err)
	assert.Equal(t, len(answers), size)

	want := score{
		points: len(problems),
		max:    len(problems),
	}

	got := runQuiz(c, problems, timer, &stdin)

	assert.Equal(t, want.points, got.points)
	assert.Equal(t, want.max, got.max)
	assert.Equal(t, want.rate(), got.rate())
}

// The functions in RunApp() are already tested. Just running them together with zero test questions.
func TestRunApp(t *testing.T) {
	fileName := "../../../test/data/problems-0of0.csv"
	timeLimit := 1

	os.Args = []string{"test", "-csv", fileName, "-limit", fmt.Sprintf("%d", timeLimit), "-shuffle"}

	RunApp()
}
