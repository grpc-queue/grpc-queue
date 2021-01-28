package server

import (
	"context"
	"fmt"

	"github.com/grpc-queue/grpc-queue/internal/queueproto"
)

type Server struct{}

func (s *Server) SayHello(context context.Context, in *queueproto.Message) (*queueproto.Message, error) {
	fmt.Println("Received: " + in.Body)
	return &queueproto.Message{Body: "hi " + in.Body}, nil
}
