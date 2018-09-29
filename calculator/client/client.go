package main

import (
	"context"
	"fmt"
	"log"
	"projects/gRPC/calculator/calculatorpb"

	"google.golang.org/grpc"
)

func findSum(c calculatorpb.CalculatorServiceClient) {
	fmt.Println("Starting to calculate sum...")

	req := &calculatorpb.CalculatorRequest{
		Numbers: &calculatorpb.Numbers{
			Num1: 3,
			Num2: 10,
		},
	}

	res, err := c.Calculator(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling Calculator RPC: %v\n", err)
	}

	log.Printf("Response from Calculator: %v", res.Sum)
}

func main() {
	fmt.Println("Hello, this is client side")

	conn, err := grpc.Dial("localhost: 50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Connection failed: %v", err)
	}

	defer conn.Close()

	c := calculatorpb.NewCalculatorServiceClient(conn)
	findSum(c)
}
