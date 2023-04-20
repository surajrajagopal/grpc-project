package main

import (
	"context"
	"log"

	"github.com/surajrajagopal/grpc-project/helloservice/proto/helloservice/pb"
	"google.golang.org/grpc"
)

func main() {
	var conn *grpc.ClientConn
	// step :1 create connect on port that server is running eg: 50051
	conn, err := grpc.Dial(":50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}

	defer conn.Close()

	// step 2: send connect to client
	c := pb.NewHelloServiceClient(conn)

	// step 3: call API (send request to API)
	resp, err := c.SayHello(context.Background(), &pb.HelloRequest{Name: "Client!"})
	if err != nil {
		log.Fatalf("Error when calling SayHello: %s", err)
	}
	log.Printf("Response from server: %s", resp.Name)
}
