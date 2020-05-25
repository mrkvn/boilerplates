package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/mrkvn/boilerplates/00/webserver/protos/hello"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type helloServer struct{}

func NewHelloServer() *helloServer {
	return &helloServer{}
}

func (*helloServer) SayHello(ctx context.Context, req *hello.HelloRequest) (*hello.HelloResponse, error) {
	fmt.Printf("SayHello was invoked with request: %v\n", req)
	name := req.GetName()
	reply := "Hello " + name
	response := hello.HelloResponse{Reply: reply}

	return &response, nil
}

func main() {
	// grpc server
	grpcServer := grpc.NewServer()
	// hello server
	helloServer := NewHelloServer()
	hello.RegisterGreeterServer(grpcServer, helloServer)

	// Register reflection service on gRPC server.
	reflection.Register(grpcServer)

	// grpc listener
	lis, err := net.Listen("tcp", "0.0.0.0:9090")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	fmt.Println("Established gRPC listener On Port 9090...")

	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatalf("Failed to serve %v", err)
	}
	// grpc connection with dgraph
	//fmt.Println("Establishing Connection with Dgraph...")
	//conn, err := grpc.Dial("localhost:9080", grpc.WithInsecure())
	//if err != nil {
	//log.Fatal(err)
	//}
	//defer conn.Close()
	//// dgraph
	//dg := dgo.NewDgraphClient(api.NewDgraphClient(conn))
	//// dgraph query
	//resp, err := dg.NewTxn().Query(context.Background(), `
	//{
	//persons(func: has(name)){
	//name
	//age
	//}
	//}
	//`)
	//if err != nil {
	//log.Fatal(err)
	//}
	//fmt.Printf("Response: %s\n", resp.Json)
}
