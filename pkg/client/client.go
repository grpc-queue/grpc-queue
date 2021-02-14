package client

import (
	"context"
	"log"

	"github.com/grpc-queue/grpc-queue/pkg/grpc/v1/queue"
	"google.golang.org/grpc"
)

type client struct {
	address string
}

func NewClient(address string) *client {
	return &client{address: address}
}

func (c *client) CreateStream(req *queue.CreateStreamRequest) (resp *queue.CreateStreamResponse) {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(c.address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	service := queue.NewQueueServiceClient(conn)
	response, err := service.CreateStream(context.Background(), req)
	if err != nil {
		log.Fatalf("Error when calling CreateStream: %s", err)
	}
	log.Printf("Response from server: %+v", response)
	return response
}
func (c *client) GetStreams(req *queue.GetStreamsRequest) (resp *queue.GetStreamsResponse) {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(c.address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	service := queue.NewQueueServiceClient(conn)
	response, err := service.GetStreams(context.Background(), req)
	if err != nil {
		log.Fatalf("Error when calling GetStreams: %s", err)
	}
	log.Printf("Response from server: %+v", response)
	return response
}
func (c *client) Push(req *queue.PushItemRequest) (resp *queue.PushItemResponse) {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(c.address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	service := queue.NewQueueServiceClient(conn)
	response, err := service.Push(context.Background(), req)
	if err != nil {
		log.Fatalf("Error when calling Push: %s", err)
	}
	log.Printf("Response from server: %+v", response)
	return response
}
func (c *client) Pop(req *queue.PopItemRequest) (resp *queue.PopItemResponse) {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(c.address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	service := queue.NewQueueServiceClient(conn)
	response, err := service.Pop(context.Background(), req)
	if err != nil {
		log.Fatalf("Error when calling Pop: %s", err)
	}
	log.Printf("Response from server: %+v", response)
	return response
}
