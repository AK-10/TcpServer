# first, use this command only once: $ docker build ./ --tag=golang_tcp_server
# if you want to run, use this command: $ docker run -p 8080:8080 -it ${image_name} -v .:/go/src/github.com/AK-10/tcpServer /bin/bash

# Use ubuntu16.04
FROM ubuntu:16.04

RUN apt-get update && apt-get upgrade -y sudo

# RUN apt-get install gcc make
RUN apt-get install -y golang-go
RUN export PATH="/usr/lib/go-1.10/bin:$PATH"
ENV GOPATH="/go"

RUN mkdir /go
RUN mkdir /go/src
RUN mkdir /go/src/github.com
RUN mkdir /go/src/github.com/AK-10
RUN mkdir /go/src/github.com/AK-10/tcpServer

# Set the working directory to /app
WORKDIR /go/src/github.com/AK-10/tcpServer
# Copy the current directory contents into the container at /app
COPY . /app
VOLUME /app

RUN apt-get update && apt-get install -y wget


