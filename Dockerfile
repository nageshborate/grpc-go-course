FROM golang
RUN apt update && apt install unzip

ENV PROTOC_VERSION 3.12.2
ENV PROTOC_GEN_GO_VERSION v1.4.2

ADD install.sh install.sh

RUN mkdir -p /go/src/github.com/nageshborate/grpc-go-course
WORKDIR /go/src/github.com/nageshborate/grpc-go-course

ENV CODE_DIR /go/src/github.com/nageshborate/grpc-go-course

ENTRYPOINT sh install.sh