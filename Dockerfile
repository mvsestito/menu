FROM golang:latest

ADD . /go/src/github.com/mvsestito/menu-api
WORKDIR /go/src/github.com/mvsestito/menu-api

RUN go get -x ./...
