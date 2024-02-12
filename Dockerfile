FROM golang:1.22-alpine

RUN apk add gcc build-base \
    apparmor-utils libapparmor libapparmor-dev

ADD . /work
WORKDIR /work/tests/e2e

RUN go mod download

RUN go build -tags apparmor \
    -ldflags '-s -w -extldflags "-static"' \
    -o /work/build/e2e main.go

# E2E tests must be running as root, as it also verifies
# hostop privileges.
USER root
ENTRYPOINT [ "/work/build/e2e" ]
