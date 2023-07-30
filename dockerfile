# syntax=docker/dockerfile:1

FROM golang:alpine AS builder

WORKDIR /app

RUN apk add g++ && apk add make

COPY . .

RUN go mod download
RUN  make build


FROM alpine:latest

COPY --from=builder /app/build /

ENV SHELL=/bin/sh

CMD ["/app"]
