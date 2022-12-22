FROM golang:1.19.3-alpine3.16


RUN mkdir /app

ADD server.go /app

WORKDIR /app

/bin/sh -c go install github.com/aws/aws-sdk-go/aws'

RUN go build -o main .



EXPOSE 8089



CMD ["go", "run", "server.go"]