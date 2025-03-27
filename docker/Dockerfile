## Stage 1: Build

FROM golang:1.22.3-alpine as builder

WORKDIR /app

COPY . ./app

RUN go mod download 

# Build the application

FROM golang:1.22.3-alpine as runner

COPY --from=builder ./app  ./app

RUN go get github.com/githubnemo/CompileDaemon

WORKDIR /app

EXPOSE 5000

CMD CompileDaemon --build="go build -o main ./app/main.go" --command=./main



