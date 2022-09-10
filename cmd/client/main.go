package main

import (
	"log"

	"github.com/o-sokol-o/grpc_log_server/pkg/client"
)

func main() {
	// Set up a connection to the server.
	target := "localhost:9000"
	grpc, err := client.New(target)
	if err != nil {
		log.Fatalf("server: %s\n%v\n", target, err)
	}
	defer grpc.Close()

	grpc.Send(0, 0, 0)
	grpc.Send(1, 1, 1)
	grpc.Send(2, 2, 2)
	grpc.Send(3, 3, 3)
	grpc.Send(4, 4, 4)
}
