## Stage 1: Build

FROM golang:1.23-alpine as builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o main ./cmd/api/main.go

# Build the application 

FROM golang:1.23-alpine as runner

WORKDIR /app

COPY --from=builder /app/main .

EXPOSE 4000

CMD ["./main"]


