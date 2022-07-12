.PHONY: all test dist

VERSION ?= dev
ifdef GITHUB_REF_NAME
VERSION = $(GITHUB_REF_NAME)
endif


default: run

all: test

build:
	go build .

demo: build
	./go-sudoku generate
	./go-sudoku -type jigsaw generate
	./go-sudoku -type samurai generate

demo-solve: build
	./go-sudoku -format csv generate | ./go-sudoku -progress solve

dist:
	gox -ldflags="-s -w -X main.version=${VERSION}" \
	    -os="linux darwin windows" \
	    -arch="amd64" \
	    -output="./dist/{{.Dir}}_{{.OS}}_{{.Arch}}" \
	    .

fmt: tidy
	go fmt $(shell go list ./...)

gen: gen-readme

gen-readme: build
	./scripts/generate_readme.sh > README.md

get-tools:
	go install github.com/mitchellh/gox@v1.0.1
	go install golang.org/x/lint/golint@latest

lint:
	golint -set_exit_status $(shell go list ./...)

run:
	go run .

test: gen fmt lint vet build
	go test -cover -coverprofile=.coverprofile $(shell go list ./...)

tidy:
	go mod tidy

vet:
	go vet $(shell go list ./...)

