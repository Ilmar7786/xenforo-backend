FROM golang:alpine as builder
WORKDIR /build

COPY go.mod .
COPY go.sum .
RUN  go mod download

RUN apk add make
COPY . .
RUN make build

ENTRYPOINT ["./build/app"]
