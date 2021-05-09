package handlers

import (
	"context"
	"log"

	pb "github.com/jegacs/multiplier-service/protos"
	"github.com/jegacs/multiplier-service/server/services"
)

type GRPCServer struct {
	pb.UnimplementedMultiplierServer
}

func (s *GRPCServer) Multiply(ctx context.Context, in *pb.MultiplierRequest) (*pb.MultiplierResponse, error) {
	log.Printf("Received: first %v, second %v", in.GetFirst(), in.GetSecond())
	service := services.NewMultiplierService(in.GetFirst(), in.GetSecond())
	result, err := service.Calculate()

	log.Printf("Result: %v", result)
	if err != nil {
		log.Printf("error: %v", err)
	}

	return &pb.MultiplierResponse{Result: result}, err
}
