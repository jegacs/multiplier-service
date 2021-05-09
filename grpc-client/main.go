package main

import (
	"context"
	"log"
	"os"
	"time"

	pb "github.com/jegacs/multiplier-service/protos"
	"google.golang.org/grpc"
)

const (
	GRPC_ADDRESS = ":9999"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(GRPC_ADDRESS, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewMultiplierClient(conn)

	// Contact the server and print out its response.
	first := ""
	second := ""
	if len(os.Args) > 1 {
		first = os.Args[1]
		second = os.Args[2]
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.Multiply(ctx, &pb.MultiplierRequest{First: first, Second: second})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Result: %s, Error: %s", r.GetResult(), r.GetError())
}
