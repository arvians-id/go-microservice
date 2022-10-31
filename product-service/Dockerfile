FROM golang:1.17-alpine

WORKDIR /app

COPY . .

RUN go build -o /product-service/cmd/server

EXPOSE 8080

CMD ./product-service/cmd/server