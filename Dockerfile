FROM golang:latest

WORKDIR /app
ENV GOPATH=/

COPY ./ /app

RUN go mod download

RUN go get github.com/githubnemo/CompileDaemon

RUN go install github.com/githubnemo/CompileDaemon

RUN go build -o ./build/scripts-launcher cmd/main.go

ENTRYPOINT CompileDaemon --build="go build -o build/scripts-launcher cmd/main.go" --command=./build/scripts-launcher