default: update
.PHONY: default

format:
	go fmt ./... && goimports -l -w . && go vet ./...
	prettier --write "**/*.md"
.PHONY: format

gen:
	./bin/algo-go update
.PHONY: gen

update: gen format
.PHONY: update

test:
	go test ./...
.PHONY: test

build:
	go build -o ./bin/algo-go ./cmd/main.go
.PHONY: build
