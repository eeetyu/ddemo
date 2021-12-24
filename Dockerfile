FROM golang:1.16.6-alpine

RUN apk update && apk add git

WORKDIR /code

COPY . .