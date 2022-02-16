.PHONY: all test dist

TOOLS := github.com/mitchellh/gox@v1.0.1 \
            golang.org/x/lint/golint \

VERSION ?= dev
ifdef GITHUB_REF_NAME
VERSION = $(GITHUB_REF_NAME)
endif


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

