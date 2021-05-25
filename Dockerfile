FROM alpine:latest

RUN mkdir /app
WORKDIR /app
ADD go-qrcode /app
CMD ["chmod","+x","./go-qrcode"]
CMD ["./go-qrcode"]