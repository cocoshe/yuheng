# syntax=docker/dockerfile:1
FROM golang:1.17-alpine
MAINTAINER cocoshe
WORKDIR /yuheng
COPY . .
RUN export GOPROXY=https://goproxy.io && go mod tidy

RUN go build

EXPOSE 8080

#CMD["./backend"]