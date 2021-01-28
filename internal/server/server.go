package server

import (
	"context"
	"errors"
	"fmt"

	"github.com/grpc-queue/grpc-queue/internal/queueproto"
)

type Server struct{}

func (s *Server) SayHello(context context.Context, in *queueproto.Message) (*queueproto.Message, error) {
	fmt.Println("Received: " + in.Body)
	return &queueproto.Message{Body: "hi " + in.Body}, nil
}

func (s *Server) CreateStream(context context.Context, in *queueproto.CreateStreamMessage) (*queueproto.ResponseMessage, error) {
	return nil, errors.New("CreateStream not implemented yet")
}
func (s *Server) DetailStream(context context.Context, in *queueproto.DetailStreamMessage) (*queueproto.DetailStreamResponseMessage, error) {
	return nil, errors.New("DetailStream not implemented yet")
}
func (s *Server) Enqueue(context context.Context, in *queueproto.EnqueueMessage) (*queueproto.ResponseMessage, error) {
	return nil, errors.New("Enqueue not implemented yet")
}
func (s *Server) Dequeue(context context.Context, in *queueproto.DequeueMessage) (*queueproto.DequeueResponseMessage, error) {
	return nil, errors.New("Dequeue not implemented yet")
}
