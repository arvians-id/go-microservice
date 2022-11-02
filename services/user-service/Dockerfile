FROM golang:1.17-alpine

WORKDIR /app

COPY . .

RUN go build -o user-service

EXPOSE 50053

CMD ./user-service