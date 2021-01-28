package client

import (
	"context"
	"log"
	"time"

	"github.com/grpc-queue/grpc-queue/internal/queueproto"
	"google.golang.org/grpc"
)

type client struct {
	address string
}

func NewClient(address string) *client {
	return &client{address: address}
}
func (c *client) Send(message string) *queueproto.Message {
	conn, err := grpc.Dial(c.address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()
	grpcClient := queueproto.NewQueueServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Hour)
	defer cancel()
	resp, err := grpcClient.SayHello(ctx, &queueproto.Message{Body: message})
	return resp
}
