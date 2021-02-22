package main

import (
	"log"
	"net"
	"strconv"

	"github.com/grpc-queue/grpc-queue/internal/server"
	"github.com/grpc-queue/grpc-queue/pkg/grpc/v1/queue"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = 9000
)

func main() {

	lis, err := net.Listen("tcp", ":"+strconv.Itoa(port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)
	queue.RegisterQueueServiceServer(grpcServer, &server.Server{})

	log.Printf("grpc server started listening on port %d", port)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}

}
