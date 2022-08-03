package app

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"log"
	math_rand "math/rand"
	"os"
	"testing"
)

/*
  Stdout testing code borrowed from Jon Calhoun's FizzBuzz example.
  https://courses.calhoun.io/lessons/les_algo_m01_08
  https://github.com/joncalhoun/algorithmswithgo.com/blob/master/module01/fizz_buzz_test.go
*/

func TestExit(t *testing.T) {
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

	// It's a silly test but I need the practice mocking.
	code := 123
	msg := "testing Exit()"
	want := fmt.Sprintf("%s\nCalling os.Exit(%d)...\n", msg, code)
	OSExitBackup := OSExit
	OSExit = func(code int) { fmt.Printf("Calling os.Exit(%d)...\n", code) }
	Exit(code, msg)
	OSExit = OSExitBackup

	// Stop capturing stdout.
	writer.Close()

	var buf bytes.Buffer
	io.Copy(&buf, testStdout)
	got := buf.String()
	if got != want {
		t.Errorf("Exit(); want %q, got %q", want, got)
	}
}

func TestShowVersion(t *testing.T) {
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
	want := "\n"
	want += fmt.Sprintf("%s Version: %s\n", metadata.name, metadata.version)
	want += fmt.Sprintf("git version: %s\n", metadata.gitVersion)
	want += fmt.Sprintf("   git hash: %s\n", metadata.gitHash)
	want += fmt.Sprintf(" build time: %s\n", metadata.buildTime)
	want += "\n"

	// Run the function who's output we want to capture.
	showVersion()

	// Stop capturing stdout.
	writer.Close()

	var buf bytes.Buffer
	io.Copy(&buf, testStdout)
	got := buf.String()
	if got != want {
		t.Errorf("showBanner(); want %q, got %q", want, got)
	}
}

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
	io.Copy(&buf, testStdout)
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
	io.Copy(&buf, testStdout)
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
		fileName: "../../../test/data/problems.csv",
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

func TestParseLines(t *testing.T) {
	c := config{
		fileName: "../../../test/data/problems.csv",
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

func TestInitRandSeed(t *testing.T) {
	// This is a difficult one, now I'm just checking to see if I can reset
	// the seed and get the same number twice. `go test` seems to be caching
	// results.
	InitRandSeed(1, true)
	want := math_rand.Intn(100)

	InitRandSeed(1, true)
	got := math_rand.Intn(100)

	if want != got {
		t.Errorf("n1, %d, isn't equal to n2, %d", want, got)
	}
}

func TestSetVersion(t *testing.T) {
	want := appInfo{
		name:       "name",
		version:    "version",
		gitVersion: "gitVersion",
		gitHash:    "gitHash",
		buildTime:  "buildTime",
	}

	SetVersion(want.version, want.gitVersion, want.gitHash, want.buildTime)

	got := metadata

	if want.version != got.version {
		t.Errorf("expected version to be set to %q, got %q", want.version, got.version)
	}

	if want.gitVersion != got.gitVersion {
		t.Errorf("expected gitVersion to be set to %q, got %q", want.gitVersion, got.gitVersion)
	}

	if want.gitHash != got.gitHash {
		t.Errorf("expected gitHash to be set to %q, got %q", want.gitHash, got.gitHash)
	}

	if want.buildTime != got.buildTime {
		t.Errorf("expected buildTime to be set to %q, got %q", want.buildTime, got.buildTime)
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
	io.Copy(&buf, testStdout)
	got := buf.String()
	if got != want {
		t.Errorf("Exit(); want %q, got %q", want, got)
	}

	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError) // flags are now reset
}
