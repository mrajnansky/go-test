FROM golang:1.25rc2-bookworm

WORKDIR /usr/src/app

RUN go install github.com/air-verse/air@latest

COPY . .
RUN go mod tidy