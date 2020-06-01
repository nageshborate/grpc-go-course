package main

import (
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	"github.com/nageshborate/grpc-go-course/greet/greetpb"
)

type server struct {

}

func main() {
	fmt.Println("Hello World")

	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer();
	greetpb.RegisterGreetServiceServer(grpcServer, &server{})

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
