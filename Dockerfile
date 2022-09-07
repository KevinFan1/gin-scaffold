FROM golang:1.18-alpine as builder

ARG PROJECT_NAME=app
ARG FOLDER=/app

WORKDIR $FOLDER
ADD . $FOLDER

ENV GO111MODULE=on \
        CGO_ENABLED=0 \
        GOOS=linux \
        GOARCH=amd64 \
    	GOPROXY="https://goproxy.io,direct"

RUN go mod tidy && go build -o $PROJECT_NAME .

FROM alpine:latest

WORKDIR $FOLDER

COPY ./config/*.yml ./config/
COPY gin-scaffold .

EXPOSE 8080

CMD ["./$PROJECT_NAME"]