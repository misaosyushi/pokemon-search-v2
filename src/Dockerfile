FROM golang:1.14

RUN mkdir /app
COPY  . /app
WORKDIR /app

RUN go build main.go
ENTRYPOINT ["./main"]

EXPOSE 8081
