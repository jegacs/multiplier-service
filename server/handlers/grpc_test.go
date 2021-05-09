package handlers

import (
	"context"
	"testing"

	pb "github.com/jegacs/multiplier-service/protos"
)

func TestGRPCMultiplyHandler(t *testing.T) {
	server := &GRPCServer{}
	t.Run("when requesting to service", func(t *testing.T) {
		t.Run("given number formats right", func(t *testing.T) {
			response, err := server.Multiply(context.Background(), &pb.MultiplierRequest{
				First:  "10",
				Second: "10",
			})

			t.Run("error should be nil", func(t *testing.T) {
				if err != nil {
					t.Errorf("error should be nil, but it was %s", err)
				}
			})

			t.Run("result should be 100.00", func(t *testing.T) {
				if response.Result != "100.00" {
					t.Errorf("result should be 100.00, but it was %s", response.Result)
				}
			})

			t.Run("Error field should be empty", func(t *testing.T) {
				if response.Error != "" {
					t.Errorf("Error field should be empty, but it was %s", response.Error)
				}
			})
		})
	})
}
