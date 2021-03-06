package handlers

import (
	"context"
	"log"
	"net"

	"github.com/jegacs/multiplier-service/errors"
	pb "github.com/jegacs/multiplier-service/protos"
	"github.com/jegacs/multiplier-service/server/services"
	"google.golang.org/grpc"
)

type GRPCServer struct {
	pb.UnimplementedMultiplierServer
}

func (s *GRPCServer) Multiply(ctx context.Context, in *pb.MultiplierRequest) (*pb.MultiplierResponse, error) {
	log.Printf("Received: first %v, second %v", in.GetFirst(), in.GetSecond())
	response := &pb.MultiplierResponse{}

	if in.GetFirst() == "" || in.GetSecond() == "" {
		response.Error = errors.ErrBadFormatNumber.Error()
		return response, errors.ErrBadFormatNumber
	}
	service := services.NewMultiplierService(in.GetFirst(), in.GetSecond())
	result, err := service.Calculate()

	log.Printf("Result: %v", result)
	response.Result = result
	if err != nil {
		log.Printf("error: %v", err)
		response.Error = err.Error()
	}
	return response, err
}

func RunGRPCServer(addr string) {
	server := &GRPCServer{}
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterMultiplierServer(s, server)
	log.Printf("Starting server in %s", addr)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
