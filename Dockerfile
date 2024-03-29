FROM golang:latest

WORKDIR /app

COPY ./ /app

RUN go mod tidy

RUN go mod download

RUN go get github.com/githubnemo/CompileDaemon

ENTRYPOINT CompileDaemon --build="go build cmd/main.go" --command="./main"