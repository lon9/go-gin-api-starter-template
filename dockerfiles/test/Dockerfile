FROM golang:alpine

RUN apk add --no-cache git gcc musl-dev
ADD . /go/src/{{ .ProjectPath }}
WORKDIR /go/src/{{ .ProjectPath }}

RUN go get
