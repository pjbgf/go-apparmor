GO ?= go
GCC ?= gcc
DOCKER ?= docker
IMAGE_TAG ?= paulinhu/go-apparmor:1
PROFILE_PATH ?= $(realpath ./example/profiles/test-profile.aa)

OUTDIR := build

LDFLAGS := -s -w -extldflags "-static"
BINARY := go-apparmor
GOSEC := gosec

.PHONY: image
image:
	$(DOCKER) build -t $(IMAGE_TAG) .

.PHONY: build
build:
	$(GO) build -ldflags '$(LDFLAGS)' -o $(OUTDIR)/$(BINARY) ./example/code/main.go

run: build
	sudo $(OUTDIR)/$(BINARY) $(PROFILE_PATH)

run-container:
	docker run --rm -it --privileged --pid host $(IMAGE_TAG) /app/go-apparmor $(PROFILE_PATH)

load-profile:
	sudo apparmor_parser -R $(PROFILE_PATH) | true
	sudo apparmor_parser -Kr $(PROFILE_PATH)
	sudo grep test-profile /sys/kernel/security/apparmor/profiles

verify:
	$(GOSEC) ./...
