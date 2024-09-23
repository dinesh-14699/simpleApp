FROM golang:1.18-alpine

WORKDIR /my_gin_app

COPY . .

RUN go build -o app .

EXPOSE 8080

CMD ["./app"]
