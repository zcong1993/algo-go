default: update
.PHONY: default

format:
	go fmt ./... && goimports -l -w . && go vet ./...
	prettier --write "**/*.md"
.PHONY: format

gen:
	go run cmd/main.go
.PHONY: gen

update: gen format
.PHONY: update

test:
	go test ./...
.PHONY: test
