package service

import (
	"context"

	pb "grpc-gateway-exp/proto/api"
)

type HelloService struct {
}

func (h *HelloService) Hello(ctx context.Context, message *pb.HelloMessage) (*pb.HelloResponse, error) {
	helloMessage := "Hello " + message.GetMessage()

	response := pb.HelloResponse{Result: helloMessage}

	return &response, nil
}
