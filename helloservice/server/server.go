package main

import (
	"fmt"
	"log"
	"net"

	"github.com/surajrajagopal/grpc-project/helloservice/configs"
	"github.com/surajrajagopal/grpc-project/helloservice/helloservice"
	"github.com/surajrajagopal/grpc-project/helloservice/proto/helloservice/pb"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Go gRPC Beginners Tutorial!")
	conf := configs.Config{
		Port: 50051,
	}

	//step 1: Listen
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", conf.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// step 2: struct imported
	s := helloservice.Server{}

	// step 3: create grpc server
	grpcServer := grpc.NewServer()

	// step 4: Register Struct and Grpc Server
	pb.RegisterHelloServiceServer(grpcServer, &s)

	// step 5: serve
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
