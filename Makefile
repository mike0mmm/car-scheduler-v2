BIN_DIR := .tools/bin

PROTOC_VERSION := 3.8.0
RELEASE_OS :=
PROTOC_DIR := $(BIN_DIR)/protoc-$(PROTOC_VERSION)
PROTOC_BIN := $(PROTOC_DIR)/bin/protoc

GO := go
ifdef GO_BIN
	GO = $(GO_BIN)
endif

GOLANGCI_LINT_VERSION := 1.27.0
GOLANGCI_LINT := $(BIN_DIR)/golangci-lint_$(GOLANGCI_LINT_VERSION)


GIT_COMMIT := $(shell git rev-parse --short HEAD 2> /dev/null || echo "no-revision")
GIT_COMMIT_MESSAGE := $(shell git show -s --format='%s' 2> /dev/null | tr ' ' _ | tr -d "'")
GIT_TAG := $(shell git describe --tags 2> /dev/null || echo "no-tag")
GIT_BRANCH := $(shell git rev-parse --abbrev-ref HEAD 2> /dev/null || echo "no-branch")
BUILD_TIME := $(shell date +%FT%T%z)

UNAME_S := $(shell uname -s)
ifeq ($(UNAME_S),Linux)
	RELEASE_OS = linux
endif
ifeq ($(UNAME_S),Darwin)
	RELEASE_OS = osx
endif

## build: build all files
build: 
	$(GO) build ./...

update:
	$(GO) get -u ./...

## lint: lint all go code
lint: $(GOLANGCI_LINT)
	$(GOLANGCI_LINT) run --fast --enable-all -D wsl -D testpackage -D godot -D goerr113

$(GOLANGCI_LINT):
	curl -sfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(BIN_DIR) v$(GOLANGCI_LINT_VERSION)
	mv $(BIN_DIR)/golangci-lint $(GOLANGCI_LINT)

run: build
	$(GO) run cmd/server/main.go

help:
	@echo "Usage: \n"
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'
