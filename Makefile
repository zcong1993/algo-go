ROOT_DIR := $(shell dirname $(realpath $(firstword $(MAKEFILE_LIST))))

# GOBIN > GOPATH > INSTALLDIR
GOBIN	:=	$(shell echo ${GOBIN} | cut -d':' -f1)
GOPATH	:=	$(shell echo $(GOPATH) | cut -d':' -f1)
BIN		:= 	""

# check GOBIN
ifneq ($(GOBIN),)
	BIN=$(GOBIN)
else
	# check GOPATH
	ifneq ($(GOPATH),)
		BIN=$(GOPATH)/bin
	endif
endif

# golangci-lint
GOLANGCI_LINT := bin/golangci-lint

$(GOLANGCI_LINT):
	curl -SL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s v1.50.0

# goimports
GOIMPORTS := bin/goimports

$(GOIMPORTS):
	GOBIN=$(ROOT_DIR)/bin go install golang.org/x/tools/cmd/goimports@latest

# lefthook
LEFTHOOK := bin/lefthook

$(LEFTHOOK):
	GOBIN=$(ROOT_DIR)/bin go install github.com/evilmartians/lefthook@latest

setup: $(GOLANGCI_LINT) $(GOIMPORTS) $(LEFTHOOK)
	$(LEFTHOOK) install
.PHONY: setup

lint: $(GOLANGCI_LINT)
	$(GOLANGCI_LINT) run --fast --fix
.PHONY: lint

default: update
.PHONY: default

gfmt:
	$(GOIMPORTS) -l -w . && go vet ./...
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
