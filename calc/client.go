package main

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"github.com/nageshborate/grpc-go-course/calc/calcpb"
)

func main() {
	clientConnection, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect %v", err)
	}
	defer clientConnection.Close()

	calculatorServiceClient := calcpb.NewCalculatorServiceClient(clientConnection)

	additionResponse, err := calculatorServiceClient.Add(context.Background(), &calcpb.AdditionRequest{
		Number1: 3,
		Number2: 10,
	})
	if err != nil {
		log.Fatalf("exception %v", err)
	}
	log.Printf("response: %v", additionResponse.Result)
}