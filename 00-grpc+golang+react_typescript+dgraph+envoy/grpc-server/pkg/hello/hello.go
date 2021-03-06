package hello

import (
	"context"
	"fmt"

	hellopb "github.com/mrkvn/boilerplates/00/grpc-server/api/protobuf/hello"
)

type helloServer struct{}

func NewHelloServer() *helloServer {
	return &helloServer{}
}

func (*helloServer) SayHello(ctx context.Context, req *hellopb.HelloRequest) (*hellopb.HelloResponse, error) {
	fmt.Printf("SayHello was invoked with request: %v\n", req)
	name := req.GetName()
	reply := "Hello " + name
	response := hellopb.HelloResponse{Reply: reply}

	return &response, nil
}
