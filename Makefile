default: update
.PHONY: default

gfmt:
	go fmt ./... && goimports -l -w . && go vet ./...
.PHONY: gfmt

format.new: gfmt
	changed-files -f "\.md$$" "prettier --write" | bash
.PHONY: format.new

format.all: gfmt
	prettier --write "**/*.md"
.PHONY: format.all

gen:
	leetcode-tool update
.PHONY: gen

update: gen format.new
.PHONY: update

test:
	go test ./...
.PHONY: test

build:
	go build -o ./bin/algo-go ./cmd/main.go
.PHONY: build

test.new:
	 changed-files --folder -f "\.go$$" "go test -v -coverprofile=coverage.txt -covermode=atomic" | bash
 .PHONY: test.new
