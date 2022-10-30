FROM golang:1.17-alpine

WORKDIR /app

COPY . .

RUN go build -o auth-service

EXPOSE 8080

CMD ./auth-service