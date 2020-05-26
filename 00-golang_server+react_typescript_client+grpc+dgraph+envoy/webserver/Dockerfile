FROM golang:1.14
MAINTAINER mrkvn

WORKDIR /app
COPY . .

RUN go mod download
RUN go get github.com/githubnemo/CompileDaemon

ENTRYPOINT CompileDaemon --build="go build cmd/webserver.go" --command=./webserver
