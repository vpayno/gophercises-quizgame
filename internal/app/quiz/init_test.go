package app

import (
	"bytes"
	"fmt"
	"io"
	math_rand "math/rand"
	"os"
	"strings"
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
	_, err = io.Copy(&buf, testStdout)
	if err != nil {
		t.Error(fmt.Sprint(err))
	}
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
	_, err = io.Copy(&buf, testStdout)
	if err != nil {
		t.Error(fmt.Sprint(err))
	}
	got := buf.String()
	if got != want {
		t.Errorf("showBanner(); want %q, got %q", want, got)
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

	strSlice := []string{want.version, want.gitVersion, want.gitHash, want.buildTime}
	bytes := []byte(strings.Join(strSlice, "\n") + "\n")
	SetVersion(bytes)

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
