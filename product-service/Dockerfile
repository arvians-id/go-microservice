FROM golang:1.17-alpine

WORKDIR /app

COPY ./cmd/server/main.go .
COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o main .

EXPOSE 50052

CMD ["./main"]