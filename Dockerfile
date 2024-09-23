FROM golang:1.18-alpine

WORKDIR /my_gin_app

COPY . .

RUN go build -o my_gin_app .

EXPOSE 8080

CMD ["./my_gin_app"]
