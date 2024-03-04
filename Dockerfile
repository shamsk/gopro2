FROM golang:1.19.0

WORKDIR  /usr/src/app

COPY . .

RUN go install github.com/cosmtrek/air@latest

RUN go mod tidy

