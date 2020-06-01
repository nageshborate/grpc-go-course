package main

import (
	"fmt"
	"google.golang.org/grpc"
	"log"
	"github.com/nageshborate/grpc-go-course/greet/greetpb"
)

func main() {
	clientConnection, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect %v", err)
	}
	defer clientConnection.Close()

	greetpb.NewGreetServiceClient(clientConnection)

	fmt.Println("created client %f", clientConnection)
}
