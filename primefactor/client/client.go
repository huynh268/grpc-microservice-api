package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"projects/gRPC/primefactor/primefactorpb"

	"google.golang.org/grpc"
)

func createServerStreaming(n int32, c primefactorpb.PrimeFactorServiceClient) {
	fmt.Println("Starting to create Server Streaming RPC...")

	req := &primefactorpb.PrimeFactorRequest{
		Number: &primefactorpb.Number{
			Num: n,
		},
	}

	resStream, err := c.PrimeFactor(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling PrimeFactor %v", err)
	}

	for {
		msg, err := resStream.Recv()
		if err == io.EOF {
			// Reach the end of streaming
			break
		}

		if err != nil {
			log.Fatalf("Error while reading streaming %v", err)
		}

		log.Printf("Response from PrimeFactor: %v", msg.GetPrime())
	}
}

func main() {
	fmt.Println("Hello, this is client side")
	conn, err := grpc.Dial("localhost: 50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect %v", err)
	}
	defer conn.Close()

	c := primefactorpb.NewPrimeFactorServiceClient(conn)

	nums := [5]int32{120, 210, 1500, 12345, 456210}
	for _, num := range nums {
		fmt.Printf("Number: %v\n", num)
		createServerStreaming(num, c)
	}
}
