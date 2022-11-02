FROM golang:1.17-alpine

WORKDIR /app

COPY . .

RUN go build -o auth-service

EXPOSE 50051

CMD ./auth-service