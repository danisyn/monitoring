FROM golang:1.20.4-alpine3.18

WORKDIR /app

RUN apk add nano
RUN apk add bash

COPY structures /app/structures

RUN go mod init ingress-monitor
RUN go mod tidy

RUN go get k8s.io/apimachinery/pkg/apis/meta/v1 \
&& go get k8s.io/client-go/kubernetes \
&& go get k8s.io/client-go/rest

#RUN go mod download

COPY *.go ./

RUN go build -o /ingress-monitor

CMD ["sleep", "infinity"]