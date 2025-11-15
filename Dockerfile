FROM golang:1.23-alpine AS builder

WORKDIR /app

RUN apk add git

ENV GO111MODULE=on

ENV CGO_ENABLED=1
# ENV CGO_ENABLED=0
ENV GOOS=linux

COPY go.mod .
COPY go.sum .

COPY . .

RUN apk update

# for kafka
RUN apk add librdkafka-dev gcc libc-dev make

RUN make ci && make build

FROM alpine:latest AS release
# FROM ubuntu:latest AS release

RUN apk add --no-cache --update ca-certificates curl

COPY --from=builder /app/main /app/cmd/
# COPY --from=builder /app/.env /app/
COPY --from=builder /app/templates /app/templates/

RUN chmod +x /app/cmd/main

WORKDIR /app

EXPOSE 80

CMD ["cmd/main"]