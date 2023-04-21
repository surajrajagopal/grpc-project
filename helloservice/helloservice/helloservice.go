package helloservice

import (
	"context"
	"log"

	"github.com/surajrajagopal/grpc-project/helloservice/proto/helloservice/pb"
)

type Server struct {
	pb.UnimplementedHelloServiceServer
}

func (s *Server) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Hello I am : %s", req.Name)
	resp := &pb.HelloReply{
		Name: "Hello I am server",
	}

	return resp, nil
}
