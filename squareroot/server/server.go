package main

import (
	"context"
	"fmt"
	"log"
	"math"
	"net"
	"projects/gRPC/squareroot/squarerootpb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type server struct{}

func (*server) SquareRoot(ctx context.Context, req *squarerootpb.SquareRootRequest) (*squarerootpb.SquareRootResponse, error) {
	fmt.Println("Receive SquareRoot RPC")

	number := req.GetNumber()
	if number < 0 {
		return nil, status.Errorf(
			codes.InvalidArgument,
			fmt.Sprintf("Received a negative number: %v\n", number),
		)
	}

	return &squarerootpb.SquareRootResponse{
		Root: math.Sqrt(float64(number)),
	}, nil
}

func main() {
	fmt.Println("Hello, this is server side")

	lis, err := net.Listen("tcp", "0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	squarerootpb.RegisterSquareRootServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
