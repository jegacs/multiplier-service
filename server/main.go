package main

import (
	"flag"
	"log"

	"github.com/jegacs/multiplier-service/server/handlers"
)

const (
	GRPC_ADDRESS = "localhost:9999"
	HTTP_ADDRESS = "localhost:8000"
)

func main() {
	mode := flag.String("mode", "", "Mode to run the server: http or grpc mode. gRPC mode is run if no argument is passed.")

	switch *mode {
	case "":
		handlers.RunGRPCServer(GRPC_ADDRESS)
	case "grpc":
		handlers.RunGRPCServer(GRPC_ADDRESS)
	case "http":
		handlers.RunHTTPServer(HTTP_ADDRESS)
	default:
		log.Fatalln("unknown mode, usage: -mode={grpc|http}")
	}
}
