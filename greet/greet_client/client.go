package main

import (
	"context"
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

	greetServiceClient := greetpb.NewGreetServiceClient(clientConnection)

	fmt.Println("created client %f", greetServiceClient)
	greetResponse, err := greetServiceClient.Greet(context.Background(), &greetpb.GreetingRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Test1",
			LastName:  "Test2",
		},
	})
	if err != nil {
		log.Fatalf("exception %v", err)
	}
	log.Printf("response: %v", greetResponse.Result)
}
