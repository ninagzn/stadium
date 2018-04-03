FROM golang:latest

ENV SRC_DIR=/go/src/stadium/
ADD . $SRC_DIR
WORKDIR $SRC_DIR
RUN go get github.com/gorilla/mux
RUN go build

ENTRYPOINT ./stadium
EXPOSE 8080