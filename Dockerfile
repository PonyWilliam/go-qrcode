FROM golang:alpine

MAINTAINER  "william"
RUN mkdir /app
WORKDIR /app/
ADD . /app/
RUN go env -w GO111MODULE="on"
RUN go env -w GOPROXY="https://goproxy.io"
RUN go env -w GOOS="linux"
RUN go build -o go-qrcode .
ENV GOPATH=/root/go
EXPOSE 80
ENTRYPOINT  ["./go-qrcode"]