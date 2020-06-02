package main

import (
"context"
"fmt"
"google.golang.org/grpc"
"log"
"net"
"github.com/nageshborate/grpc-go-course/calc/calcpb"
)

type server struct {}

func (*server) Add(ctx context.Context, additionRequest *calcpb.AdditionRequest) (*calcpb.AdditionResponse, error) {
	return &calcpb.AdditionResponse {
		Result: additionRequest.GetNumber1() + additionRequest.GetNumber2(),
	}, nil
}

func main() {
	fmt.Println("Starting server...")

	listener, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer();
	calcpb.RegisterCalculatorServiceServer(grpcServer, &server{})

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
