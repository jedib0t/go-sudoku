.PHONY: all test

default: run

all: test

build:
	go build .

demo-generate: build
	./go-sudoku generate
	./go-sudoku -type jigsaw generate
	./go-sudoku -type samurai generate

demo-solve: build
	./go-sudoku -format csv generate | ./go-sudoku -progress solve

fmt:
	go fmt $(shell go list ./...)

gen: gen-readme

gen-readme: build
	./scripts/generate_readme.sh > README.md

lint:
	golint -set_exit_status $(shell go list ./...)

run:
	go run .

test: gen fmt lint vet build
	go test -cover -coverprofile=.coverprofile $(shell go list ./...)

vet:
	go vet $(shell go list ./...)

