## Stage 1: Build

FROM golang:1.22.3-alpine as builder

WORKDIR /app

COPY ./ ./app

COPY go.mod go.sum ./

# List the contents of the current directory
RUN ls -la

RUN go mod download 

# Build the application

FROM golang:1.22.3-alpine as runner

COPY --from=builder ./app  ./app

RUN go install github.com/githubnemo/CompileDaemon@latest

WORKDIR /app

EXPOSE 5000

CMD CompileDaemon --build="go build -o main ./cmd/api/main.go" --command=./main



