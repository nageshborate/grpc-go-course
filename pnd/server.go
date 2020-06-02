package main

import (
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	"github.com/nageshborate/grpc-go-course/pnd/pndpb"
)

type server struct {}

func (*server) Decompose(req *pndpb.PrimeNumberDecompositionRequest, res pndpb.PrimeNumberDecompositionService_DecomposeServer) error {
	factor := int64(2)
	inputNumber := req.GetInputNumber()
	for inputNumber > 1 {
		if inputNumber % factor == 0 {
			res.Send(&pndpb.PrimeNumberDecompositionResponse {
				Result: factor,
			})
			inputNumber = inputNumber / factor
		} else {
			factor++
		}
	}
	return nil
}

func main() {
	fmt.Println("Starting server")

	listener, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer();
	pndpb.RegisterPrimeNumberDecompositionServiceServer(grpcServer, &server{})

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
