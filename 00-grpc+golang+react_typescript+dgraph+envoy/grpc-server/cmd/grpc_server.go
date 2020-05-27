package main

import (
	"fmt"
	"log"
	"net"

	dgraphpb "github.com/mrkvn/boilerplates/00/grpc-server/api/protobuf/dgraph"
	hellopb "github.com/mrkvn/boilerplates/00/grpc-server/api/protobuf/hello"
	"github.com/mrkvn/boilerplates/00/grpc-server/pkg/dgraph"
	"github.com/mrkvn/boilerplates/00/grpc-server/pkg/hello"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	grpcServer := grpc.NewServer()
	helloServer := hello.NewHelloServer()
	hellopb.RegisterGreeterServer(grpcServer, helloServer)
	dgraphServer := dgraph.NewDgraphServer()
	dgraphpb.RegisterQueryServer(grpcServer, dgraphServer)

	reflection.Register(grpcServer)

	lis, err := net.Listen("tcp", ":9090")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	fmt.Println("Established gRPC listener On Port 9090...")

	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatalf("Failed to serve %v", err)
	}
}
