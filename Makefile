# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=go-update-test
LD_FLAGS=-X 'main.version=$$(git describe)'

deps:
		export GO111MODULE=on
		export GOPROXY=direct
		export GOSUMDB=off
		$(GOGET) ./...

build:	deps
		export GO111MODULE=on
		[ -d bin ] || mkdir bin
		$(GOBUILD) -ldflags "$(LD_FLAGS)" -o bin/$(BINARY_NAME) -v .
