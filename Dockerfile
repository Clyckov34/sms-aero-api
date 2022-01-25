FROM golang:latest
RUN mkdir /app

ADD . /app/
WORKDIR /app

RUN go build -o app cmd/sms/main.go
CMD ["./app"]