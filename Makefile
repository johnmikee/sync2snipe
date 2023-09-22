all: build

.PHONY: build

ifeq ($(GOPATH),)
	PATH := $(HOME)/go/bin:$(PATH)
else
	PATH := $(GOPATH)/bin:$(PATH)
endif

export GO111MODULE=on

BRANCH = $(shell git rev-parse --abbrev-ref HEAD)
REVISION = $(shell git rev-parse HEAD)
REVSHORT = $(shell git rev-parse --short HEAD)
USER = $(shell whoami)
GOVERSION = $(shell go version | awk '{print $$3}')
NOW	= $(shell date -u +"%Y-%m-%dT%H:%M:%SZ")
SHELL = /bin/sh
VERSION = $(shell git describe --tags --always)

ifneq ($(OS), Windows_NT)
	CURRENT_PLATFORM = linux
	ifeq ($(shell uname), Darwin)
		SHELL := /bin/sh
		CURRENT_PLATFORM = darwin
	endif
else
	CURRENT_PLATFORM = windows
endif

BUILD_VERSION = "\
	-X github.com/ramp/it/sync2snipe/pkg/version.appName=${APP_NAME} \
	-X github.com/ramp/it/sync2snipe/pkg/version.version=${VERSION} \
	-X github.com/ramp/it/sync2snipe/pkg/version.branch=${BRANCH} \
	-X github.com/ramp/it/sync2snipe/pkg/version.buildUser=${USER} \
	-X github.com/ramp/it/sync2snipe/pkg/version.buildDate=${NOW} \
	-X github.com/ramp/it/sync2snipe/pkg/version.revision=${REVISION} \
	-X github.com/ramp/it/sync2snipe/pkg/version.goVersion=${GOVERSION}"

deps:
	go mod download

test:
	go test -cover ./...

build: s2s

clean:
	rm -rf build/
	rm -f *.zip

.pre-build:
	mkdir -p build/darwin
	mkdir -p build/linux

install-local: install-s2s

APP_NAME = sync2snipe

.pre-s2s:
	$(eval APP_NAME = sync2snipe)

s2s: .pre-build .pre-s2s
	go build -o build/$(CURRENT_PLATFORM)/s2s -ldflags ${BUILD_VERSION} ./cmd/s2s

install-s2s: .pre-s2s
	go install -ldflags ${BUILD_VERSION} ./cmd/s2s

docker-s2s:
	cp resources/docker/Dockerfile .
	docker build -t s2s --rm .
	rm Dockerfile

run-docker-s2s:
	docker run sync2snipe sync2snipe

static-check:
	staticcheck ./...

vet:
	go vet ./...

tidy:
	go mod tidy
