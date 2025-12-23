FROM golang:1.25-alpine AS builder

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

# Install swag for Swagger documentation
ENV GOPATH=/go
RUN go install github.com/swaggo/swag/cmd/swag@latest

# Add GOPATH/bin to PATH and generate Swagger documentation
ENV PATH="/go/bin:${PATH}"
RUN /go/bin/swag init -g cmd/main.go -o docs

RUN make ci && make build

FROM alpine:latest AS release
# FROM ubuntu:latest AS release

RUN apk add --no-cache --update ca-certificates curl

COPY --from=builder /app/main /app/cmd/
COPY --from=builder /app/docs /app/docs/

RUN chmod +x /app/cmd/main

WORKDIR /app

EXPOSE 80

CMD ["cmd/main"]