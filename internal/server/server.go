package server

import (
	"context"
	"errors"

	"github.com/grpc-queue/grpc-queue/pkg/grpc/v1/queue"
)

type Server struct {
	*queue.UnimplementedQueueServiceServer
}

func (s *Server) CreateStream(ctx context.Context, request *queue.CreateStreamRequest) (*queue.CreateStreamResponse, error) {
	return nil, errors.New("CreateStream not implemented yet")
}
func (s *Server) GetStreams(ctx context.Context, request *queue.GetStreamsRequest) (*queue.GetStreamsResponse, error) {
	return nil, errors.New("GetStreams not implemented yet")
}

func (s *Server) Push(ctx context.Context, request *queue.PushItemRequest) (*queue.PushItemResponse, error) {
	return nil, errors.New("Push not implemented yet")
}
func (s *Server) Pop(request *queue.PopItemRequest, service queue.QueueService_PopServer) error {
	return errors.New("Pop not implemented yet")
}
