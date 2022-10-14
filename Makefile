SHELL=/bin/bash

GO ?= go
GCC ?= gcc
DOCKER ?= docker
BUILD_TAGS ?= apparmor

IMAGE_TAG = quay.io/paulinhu/go-apparmor/e2e:local

CWD := $(realpath .)
OUTDIR := $(CWD)/build

LDFLAGS := -s -w -extldflags "-static"
BINARY := go-apparmor
GOSEC := gosec


.PHONY: build
build:
	$(GO) build -tags $(BUILD_TAGS) ./...

tidy:
	$(GO) mod tidy
	pushd tests/e2e && \
	$(GO) mod tidy || \
	popd

verify: tidy
	$(GOSEC) ./...

test:
	$(GO) test -tags $(BUILD_TAGS) ./...

e2e:
	$(DOCKER) build -t $(IMAGE_TAG) .
	$(DOCKER) run --rm -it --privileged --pid host $(IMAGE_TAG)
