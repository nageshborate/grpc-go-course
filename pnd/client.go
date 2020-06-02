package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"io"
	"log"
	"github.com/nageshborate/grpc-go-course/pnd/pndpb"
)

func main() {
	clientConnection, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect %v", err)
	}
	defer clientConnection.Close()

	primeNumberDecompositionServiceClient := pndpb.NewPrimeNumberDecompositionServiceClient(clientConnection)

	fmt.Println("created client %f", primeNumberDecompositionServiceClient)
	greetResponse, err := primeNumberDecompositionServiceClient.Decompose(context.Background(), &pndpb.PrimeNumberDecompositionRequest {InputNumber: 120})
	if err != nil {
		log.Fatalf("exception %v", err)
	}

	for {
		result, err := greetResponse.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("error %v", err)
		}
		log.Println("result: ", result.GetResult())
	}
}
