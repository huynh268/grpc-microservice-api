package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"projects/gRPC/calculator/calculatorpb"

	"google.golang.org/grpc"
)

type server struct{}

func (*server) Calculator(ctx context.Context, req *calculatorpb.CalculatorRequest) (*calculatorpb.CalculatorResponse, error) {
	fmt.Printf("Calculator function is invoked with %v\n", req)
	num1 := req.GetNumbers().GetNum1()
	num2 := req.GetNumbers().GetNum2()
	sum := num1 + num2

	res := &calculatorpb.CalculatorResponse{
		Sum: sum,
	}
	return res, nil
}

func main() {
	fmt.Println("Hello, This is server side")

	lis, err := net.Listen("tcp", "0.0.0:50051")
	if err != nil {
		log.Fatalf("Connection failed: %v", err)
	}

	s := grpc.NewServer()

	calculatorpb.RegisterCalculatorServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to server: %v", err)
	}
}
