FROM golang:latest

LABEL maintainer="Zach Hanson, <zacharyjhanson@gmail.com"

WORKDIR /app

COPY ./go.mod ./

RUN go mod download

COPY . .

RUN go build -o main .

EXPOSE 8080

CMD ["./main"]up