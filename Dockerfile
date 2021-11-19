FROM golang:1.17-bullseye

RUN apt update && apt install -y apparmor-utils libapparmor-dev

ADD . /work
WORKDIR /work
RUN make build



FROM gcr.io/distroless/static

COPY --from=0 /sbin/apparmor_parser /sbin
COPY --from=0 /work/build /app

ENTRYPOINT [ "/app/go-apparmor" ]
