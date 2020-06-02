package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	"github.com/nageshborate/grpc-go-course/greet/greetpb"
)

type server struct {}

func (*server) Greet(ctx context.Context, greetingRequest *greetpb.GreetingRequest) (*greetpb.GreetingResponse, error) {
	return &greetpb.GreetingResponse{
		Result: "Hey " + greetingRequest.Greeting.GetFirstName(),
	}, nil
}

func main() {
	fmt.Println("Hello World")

	listener, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer();
	greetpb.RegisterGreetServiceServer(grpcServer, &server{})

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
