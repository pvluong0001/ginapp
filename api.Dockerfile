FROM golang:1.15.6-alpine3.12

MAINTAINER pvluong0001@gmail.com

RUN apk --update --no-cache add g++ git && \
    apk del --purge g++

RUN go get -u github.com/cosmtrek/air github.com/google/wire/cmd/wire

ENV GIN_MODE=release

WORKDIR /build
RUN chmod 755 /build

RUN mkdir -p /build/tmp
RUN chown 1000:1000 /build/tmp

RUN addgroup -g 1000 -S gouser && \
    adduser -u 1000 -S gouser -G gouser

USER gouser

CMD air