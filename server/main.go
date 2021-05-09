package main

import (
	"log"
	"net"

	pb "github.com/jegacs/multiplier-service/protos"
	"github.com/jegacs/multiplier-service/server/handlers"
	"google.golang.org/grpc"
)

const (
	ADDRESS = "localhost:9999"
)

func main() {
	lis, err := net.Listen("tcp", ADDRESS)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterMultiplierServer(s, &handlers.GRPCServer{})
	log.Printf("Starting server in %s", ADDRESS)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
