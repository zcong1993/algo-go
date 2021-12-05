default: update
.PHONY: default

format:
	go fmt ./... && goimports -l -w . && go vet ./...
	changed-files -f "\.md$$" "prettier --write" | bash
.PHONY: format

gen:
	leetcode-tool update
.PHONY: gen

update: gen format
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
