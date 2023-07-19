FROM golang:1.20.4-alpine3.18

WORKDIR /app

COPY go.mod ./

RUN go mod download

COPY *.go ./

RUN go get ingress-monitor

RUN go build -o /ingress-monitor

CMD ["sleep", "infinity"]