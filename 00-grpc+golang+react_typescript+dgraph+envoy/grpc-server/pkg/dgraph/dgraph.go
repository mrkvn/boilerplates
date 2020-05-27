package dgraph

import (
	"context"
	"fmt"
	"log"

	"github.com/dgraph-io/dgo/v200"
	"github.com/dgraph-io/dgo/v200/protos/api"
	dgraphpb "github.com/mrkvn/boilerplates/00/grpc-server/api/protobuf/dgraph"
	"google.golang.org/grpc"
)

type dgraphServer struct{}

func NewDgraphServer() *dgraphServer {
	return &dgraphServer{}
}

func (*dgraphServer) Query(ctx context.Context, req *dgraphpb.QueryRequest) (*dgraphpb.QueryResponse, error) {
	fmt.Printf("DraphQuery was invoked with request: %v\n", req)
	query := req.GetQuery()
	fmt.Println("Establishing Connection with Dgraph...")

	// with docker-compose
	conn, err := grpc.Dial("dgraph-alpha:9080", grpc.WithInsecure())
	if err != nil {
		// without docker-compose
		conn, err = grpc.Dial("localhost:9080", grpc.WithInsecure())
		if err != nil {
			fmt.Printf("Cannot establish connection with Dgraph")
			log.Fatal(err)
		}
	}
	defer conn.Close()

	dg := dgo.NewDgraphClient(api.NewDgraphClient(conn))
	resp, err := dg.NewTxn().Query(context.Background(), query)
	if err != nil {
		fmt.Println("Cannot query from Dgraph")
	}
	fmt.Printf("Response: %s\n", resp.Json)
	response := dgraphpb.QueryResponse{Response: string(resp.Json)}

	return &response, nil
}
