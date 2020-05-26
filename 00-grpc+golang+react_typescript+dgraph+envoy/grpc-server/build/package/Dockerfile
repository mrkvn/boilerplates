FROM golang:1.14
MAINTAINER mrkvn

WORKDIR /app
COPY . .

RUN go mod download
RUN go get github.com/githubnemo/CompileDaemon

ENTRYPOINT CompileDaemon --build="go build cmd/grpc_server.go" --command=./grpc_server
