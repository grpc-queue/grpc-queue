package client

import (
	"context"
	"fmt"
	"io"
	"log"

	"github.com/grpc-queue/grpc-queue/pkg/grpc/v1/queue"
	"google.golang.org/grpc"
)

type client struct {
	conn *grpc.ClientConn
}

func NewClient(c *grpc.ClientConn) *client {
	return &client{conn: c}
}

func (c *client) CreateStream(ctx context.Context, req *queue.CreateStreamRequest) (resp *queue.CreateStreamResponse, err error) {
	service := queue.NewQueueServiceClient(c.conn)
	response, err := service.CreateStream(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", "Error when calling CreateStream", err)
	}
	log.Printf("Response from server: %+v", response)
	return response, nil
}
func (c *client) GetStreams(ctx context.Context, req *queue.GetStreamsRequest) (resp *queue.GetStreamsResponse, err error) {

	service := queue.NewQueueServiceClient(c.conn)
	response, err := service.GetStreams(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", "Error when calling GetStreams", err)
	}
	log.Printf("Response from server: %+v", response)
	return response, nil
}
func (c *client) Push(ctx context.Context, req *queue.PushItemRequest) (resp *queue.PushItemResponse, err error) {

	service := queue.NewQueueServiceClient(c.conn)
	response, err := service.Push(ctx, req)
	if err != nil {
		log.Fatalf("Error when calling Push: %s", err)
		return nil, fmt.Errorf("%s: %w", "Error when calling Push", err)
	}
	log.Printf("Response from server: %+v", response)
	return response, nil
}

func (c *client) Pop(ctx context.Context, req *queue.PopItemRequest, processFn func(*queue.PopItemResponse)) error {
	service := queue.NewQueueServiceClient(c.conn)
	re, err := service.Pop(ctx, req)
	if err != nil {
		return fmt.Errorf("%s: %w", "Error when calling Pop", err)
	}

	for {
		item, err := re.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Println(fmt.Errorf("%s: %w", "something went wrong on the pop stream", err))
			return err
		}
		processFn(item)
	}
	return nil
}
