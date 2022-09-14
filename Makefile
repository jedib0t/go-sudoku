.PHONY: all test dist

VERSION ?= dev
ifdef GITHUB_REF_NAME
VERSION = $(GITHUB_REF_NAME)
endif


default: run

all: test

build: build-game build-generator

build-game:
	go build -o go-sudoku ./cmd/game

build-generator:
	go build -o go-sudoku-generator ./cmd/generator

demo-generator: build-generator
	./go-sudoku-generator generate
	./go-sudoku-generator -type jigsaw generate
	./go-sudoku-generator -type samurai generate

demo-solver: build-generator
	./go-sudoku-generator -format csv generate | ./go-sudoku-generator -progress solve

dist: dist-game dist-generator

dist-game:
	gox -ldflags="-s -w -X main.version=${VERSION}" \
	    -os="linux darwin windows" \
	    -arch="amd64" \
	    -output="./dist/go-sudoku_{{.OS}}_{{.Arch}}_${VERSION}" \
	    ./cmd/game

dist-generator:
	gox -ldflags="-s -w -X main.version=${VERSION}" \
	    -os="linux darwin windows" \
	    -arch="amd64" \
	    -output="./dist/go-sudoku-generator_{{.OS}}_{{.Arch}}_${VERSION}" \
	    ./cmd/generator

fmt: tidy
	go fmt $(shell go list ./...)

gen: gen-readme

gen-readme: build
	./docs/generate_readme.sh > README.md

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

