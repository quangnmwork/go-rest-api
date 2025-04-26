## Stage 1: Build

FROM golang:1.22.3-alpine as builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download 

COPY . .

RUN go build -o /main ./cmd/api/main.go

# Build the application 

FROM golang:1.22.3-alpine as runner

COPY --from=builder /main /main

WORKDIR /app

EXPOSE 4000

CMD ["./main"]


