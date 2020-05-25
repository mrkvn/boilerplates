FROM golang:1.14
MAINTAINER mrkvn

WORKDIR $GOPATH/src/boilerplate/webserver
COPY . .

RUN go mod download
RUN go mod verify
RUN GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /go/bin/webserver
ENTRYPOINT ["/go/bin/webserver"]
