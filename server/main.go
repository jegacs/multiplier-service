package main

import (
	"github.com/jegacs/multiplier-service/server/handlers"
)

const (
	GRPC_ADDRESS = "localhost:9999"
	HTTP_ADDRESS = "localhost:8000"
)

func main() {
	handlers.RunGRPCServer(GRPC_ADDRESS)
	handlers.RunHTTPServer(HTTP_ADDRESS)
}
