FROM golang:1.17-alpine

WORKDIR /app

COPY . .

RUN go build -o /user-service/cmd/server

EXPOSE 8080

CMD ./user-service/cmd/server