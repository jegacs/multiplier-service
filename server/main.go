package main

import (
	"flag"
	"log"

	"github.com/jegacs/multiplier-service/server/handlers"
)

const (
	GRPC_ADDRESS = ":9999"
	HTTP_ADDRESS = ":8000"
)

func main() {
	httpFlag := flag.Bool("http", false, "Run the server in http mode.")
	grpcFlag := flag.Bool("grpc", true, "Run the server in grpc mode.")
	flag.Parse()

	if *httpFlag {
		log.Println("running http mode")
		handlers.RunHTTPServer(HTTP_ADDRESS)
		return
	}

	if *grpcFlag {
		log.Println("running in grpc mode")
		handlers.RunGRPCServer(GRPC_ADDRESS)
		return
	}
}
