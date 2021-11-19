SHELL=/bin/bash

GO ?= go
GCC ?= gcc
DOCKER ?= docker
IMAGE_TAG ?= paulinhu/go-apparmor:1

CWD := $(realpath .)
OUTDIR := $(CWD)/build
PROFILE_PATH ?= $(CWD)/example/profiles/test-profile.aa

LDFLAGS := -s -w -extldflags "-static"
BINARY := go-apparmor
GOSEC := gosec

.PHONY: image
image:
	$(DOCKER) build -t $(IMAGE_TAG) .

.PHONY: build
build:
	pushd example/code && \
	$(GO) build -ldflags '$(LDFLAGS)' -o $(OUTDIR)/$(BINARY) ./main.go || \
	popd

.PHONY: run
run: build
	$(OUTDIR)/$(BINARY) $(PROFILE_PATH)

.PHONY: run-container
run-container: image
	docker run --rm -it --privileged --pid host $(IMAGE_TAG) $(PROFILE_PATH)

.PHONY: load-profile
load-profile:
	apparmor_parser -R $(PROFILE_PATH) | true
	apparmor_parser -Kr $(PROFILE_PATH)
	grep test-profile /sys/kernel/security/apparmor/profiles

.PHONY: verify
verify:
	$(GOSEC) ./...
