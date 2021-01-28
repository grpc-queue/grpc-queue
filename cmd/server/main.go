package main

import (
	"log"
	"net"

	"github.com/grpc-queue/grpc-queue/internal/queueproto"
	"github.com/grpc-queue/grpc-queue/internal/server"
	"google.golang.org/grpc"
)

func main() {

	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	queueproto.RegisterQueueServiceServer(grpcServer, &server.Server{})

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}

}
