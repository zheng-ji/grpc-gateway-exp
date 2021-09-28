package server

import (
	"log"
	"net"

	"google.golang.org/grpc"
	pb "grpc-gateway-exp/proto/api"
	"grpc-gateway-exp/service"
)

var helloService = service.HelloService{}

func StartGrpcServer() {
	listener, err := net.Listen("tcp", ":9090")
	if err != nil {
		log.Fatalln("Listen gRPC port failed: ", err)
	}

	server := grpc.NewServer()
	pb.RegisterHelloServiceServer(server, &helloService)

	log.Println("Start gRPC Server on 0.0.0.0:9090")
	err = server.Serve(listener)
	if err != nil {
		log.Fatalln("Start gRPC Server failed: ", err)
	}

}
