SHELL := /bin/bash

GO_IMAGE_NAME = gcr.io/pendo-dev/go-build
GO_IMAGE_VERSION = 1.11-build-176
PROJECT = github.com/pendo-io/jzb
PWD := $(shell pwd)
GUID := $(shell id -g)

.DEFAULT_GOAL := default

clean:
		@echo "Cleaning build.."
		@rm -rf dist

build-arch:
		@echo "Building for darwin/386"
		@mkdir -p dist/darwin
		@GOOS=darwin GOARCH=386 go build -mod vendor -v -o dist/darwin/jzb cmd/jzb.go
		@echo "Building for linux/386"
		@mkdir -p dist/linux
		@GOOS=linux GOARCH=386 go build -mod vendor -v -o dist/linux/jzb cmd/jzb.go
		@mkdir dist/linux/release dist/darwin/release
		@chmod 775 -R dist/
		@tar -c -f dist/linux/release/jzb-linux.tar dist/linux/jzb
		@tar -c -f dist/darwin/release/jzb-darwin.tar dist/darwin/jzb

build-arch-in-container: .check-docker
		@docker run --user :${GUID} -e GO111MODULE=on -v ${PWD}:/go/src/${PROJECT}:Z -w /go/src/${PROJECT} ${GO_IMAGE_NAME}:${GO_IMAGE_VERSION} make build-arch

test:
		@echo "Running tests.."
		@go test -mod vendor -v ./...

test-in-container: .check-docker
		@docker run -e GO111MODULE=on -v ${PWD}:/go/src/${PROJECT}:Z -w /go/src/${PROJECT} ${GO_IMAGE_NAME}:${GO_IMAGE_VERSION} make test

check:
		@echo "Checking format.."
		@./tools/check-format

check-in-container: .check-docker
		@docker run -v ${PWD}:/go/src/${PROJECT}:Z -w /go/src/${PROJECT} ${GO_IMAGE_NAME}:${GO_IMAGE_VERSION} make check

format: .check-docker
		@echo "Formatting.."
		@docker run -v ${PWD}:/go/src/${PROJECT}:Z -w /go/src/${PROJECT} ${GO_IMAGE_NAME}:${GO_IMAGE_VERSION} gofmt -w ./cmd ./pkg ./internal

default: clean build-arch-in-container check-in-container test-in-container
		@echo "done"

.check-docker:
		@docker ps &>/dev/null || (echo "docker not running, or not installed."; exit 1)
