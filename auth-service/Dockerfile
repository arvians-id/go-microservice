FROM golang:1.17-alpine

WORKDIR /app

COPY . .

RUN go build -o /auth-service/cmd/server

EXPOSE 8080

CMD ./auth-service/cmd/server