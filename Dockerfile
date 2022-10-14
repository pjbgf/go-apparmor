FROM golang:1.19-alpine

RUN apk add gcc build-base \
    apparmor-utils libapparmor libapparmor-dev

ADD . /work
WORKDIR /work/tests/e2e

RUN go mod download

RUN go build -tags apparmor \
    -ldflags '-s -w -extldflags "-static"' \
    -o /work/build/e2e main.go

ENTRYPOINT [ "/work/build/e2e" ]
