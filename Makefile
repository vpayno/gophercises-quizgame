#
# Makefile
#

.DEFAULT_GOAL := all

.PHONY: all run build version clean prepare gotest cover annotate bench check vet ineffassign lint gocyclo gocritic golangci-lint misspell vendor tidy gosec

BIN_FILE=./build/quiz

all: clean annotate check

check: vet ineffassign lint gocyclo gocritic golangci-lint misspell

build: clean prepare
	go build ./...
	@printf "\n"

build-all:
	mkdir -pv build/
	@printf "Building for every OS and Platform\n\n"
	GOOS=freebsd GOARCH=386 go build -o "${BIN_FILE}"-freebsd-386 ./cmd/quiz/main.go
	GOOS=linux   GOARCH=386 go build -o "${BIN_FILE}"-linux-386 ./cmd/quiz/main.go
	GOOS=windows GOARCH=386 go build -o "${BIN_FILE}"-windows-386 ./cmd/quiz/main.go
	@printf "\n"

run: build prepare
	go run ./...
	@printf "\n"

version: build
	go version
	@printf "\n"

clean:
	go clean
	rm -fv .coverage.* ./cmd/quiz/.version.txt
	@printf "\n"

vendor:
	go mod vendor
	@printf "\n"

tidy:
	go mod tidy
	@printf "\n"

prepare:
	go generate ./...
	@printf "\n"

gotest: clean prepare
	gotest -v -covermode=count -coverprofile=.coverage.out -cover ./...
	@printf "\n"

cover: gotest
	gocov convert .coverage.out | gocov report
	@printf "\n"

annotate: cover
	gocov convert .coverage.out | gocov annotate -ceiling=100 -color -
	@printf "\n"

bench: clean prepare
	go test --run=xxx --bench . --benchmem  ./...
	@printf "\n"

vet: prepare
	go vet ./...
	@printf "\n"

ineffassign: prepare
	ineffassign ./...
	@printf "\n"

lint: prepare
	revive ./...
	@printf "\n"

gocyclo: prepare
	gocyclo .
	@printf "\n"

gocritic: prepare
	gocritic check -enableAll ./...
	@printf "\n"

golangci-lint: prepare
	golangci-lint run --out-format line-number ./...
	@printf "\n"

misspell: prepare
	misspell -error .
	@printf "\n"

gosec: prepare
	gosec ./...
	@printf "\n"