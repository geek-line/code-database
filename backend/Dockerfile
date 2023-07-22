FROM golang:1.18.3

WORKDIR /backend

ARG GOPATH=/go

RUN mkdir /scripts \
    && wget 'https://raw.githubusercontent.com/vishnubob/wait-for-it/master/wait-for-it.sh' -P /scripts  \
    && chmod 755 /scripts/wait-for-it.sh

RUN go install github.com/cosmtrek/air@latest

COPY go.mod go.sum /backend/
RUN go mod download
