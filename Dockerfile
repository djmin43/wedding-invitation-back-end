#!/usr/bin/env bash
FROM golang:latest

ENV PORT 80

EXPOSE 80
EXPOSE 4000
EXPOSE 8080

WORKDIR /app

# COPY go.mod ./
# COPY go.sum ./
# RUN go mod download

COPY ./marriage-invitation-back-end .

CMD ["./marriage-invitation-back-end"]

