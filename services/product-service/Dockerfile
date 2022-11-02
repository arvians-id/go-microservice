FROM golang:1.17-alpine

WORKDIR /app

COPY . .

RUN go build -o product-service

EXPOSE 50052

CMD ./product-service