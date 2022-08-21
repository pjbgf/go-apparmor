FROM golang:1.19-alpine

RUN apk add gcc build-base \
    apparmor-utils libapparmor libapparmor-dev

ADD . /work
WORKDIR /work
RUN cd tests/e2e && \
    go build -tags apparmor \
    -ldflags '-s -w -extldflags "-static"' \
    -o /work/build/e2e main.go

ENTRYPOINT [ "/work/build/e2e" ]
