FROM golang:1.20.4-alpine3.18

WORKDIR /app

RUN apk add nano
RUN apk add bash

COPY influxdb /app/influxdb

COPY structures /app/structures

COPY go.mod go.sum ./

RUN go mod download

COPY *.go ./

RUN go build -o /ingress-monitor

CMD ["sleep", "infinity"]