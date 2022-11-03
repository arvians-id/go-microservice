FROM golang:1.17-alpine

WORKDIR /app

COPY . .

RUN go build -o adapter

EXPOSE 8080

CMD ./adapter