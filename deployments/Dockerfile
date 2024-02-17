FROM golang:1.21-alpine

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod tidy

COPY . .

COPY go.mod go.sum ./

EXPOSE 8080

CMD ["./bin/konsultanku"]