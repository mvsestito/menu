FROM golang:latest
RUN apt-get update
RUN apt-get -y install postgresql
ADD . /go/src/github.com/mvsestito/menu-api
WORKDIR /go/src/github.com/mvsestito/menu-api
RUN go get -v -t ./...
RUN cd ../ && git clone https://github.com/vishnubob/wait-for-it.git
