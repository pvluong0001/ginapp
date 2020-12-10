FROM golang:1.15.6-alpine3.12

MAINTAINER pvluong0001@gmail.com

RUN apk add --no-cache git
RUN go get -u github.com/cosmtrek/air

ENV GIN_MODE=release

WORKDIR /build
RUN chmod 755 /build

RUN mkdir -p /build/tmp
RUN chown 1000:1000 /build/tmp

CMD air