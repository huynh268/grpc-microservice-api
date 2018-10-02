package main

import (
	"context"
	"fmt"
	"log"
	"projects/gRPC/squareroot/squarerootpb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func createErrorUnary(c squarerootpb.SquareRootServiceClient) {
	fmt.Println("Starting to create SquareRoot Unary RPC... ")

	// correct call
	createErrorCall(c, 10)

	// Error call
	createErrorCall(c, -2)
}

func createErrorCall(c squarerootpb.SquareRootServiceClient, n int32) {
	res, err := c.SquareRoot(context.Background(), &squarerootpb.SquareRootRequest{Number: n})
	if err != nil {
		respError, ok := status.FromError(err)
		if ok {
			// Actual error from gRPC (user error)
			fmt.Printf("Error message from server: %v", respError.Message())
			fmt.Println(respError.Code())
			if respError.Code() == codes.InvalidArgument {
				fmt.Println("Probably a negative number is sent!")
			}
		} else {
			log.Fatalf("Big Error calling SquareRoot: %v", err)
		}
	}
	fmt.Printf("Result of square root of %v: %v\n", n, res.GetRoot())
}

func main() {
	fmt.Println("Hello, this is client side")

	conn, err := grpc.Dial("localhost: 50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	c := squarerootpb.NewSquareRootServiceClient(conn)
	createErrorUnary(c)
}
