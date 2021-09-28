package main

import (
	"grpc-gateway-exp/server"
)

func main() {
	go server.StartGrpcServer()
	server.StartGwServer()
}
