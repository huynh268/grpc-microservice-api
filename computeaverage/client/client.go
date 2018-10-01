package main

import (
	"context"
	"fmt"
	"log"
	"projects/gRPC/computeaverage/computeaveragepb"
	"time"

	"google.golang.org/grpc"
)

func createClientStreaming(nums []float32, c computeaveragepb.ComputeAverageServiceClient) {
	fmt.Println("Starting to create Client Streaming RPC...")

	var requests []*computeaveragepb.ComputeAverageRequest

	for _, num := range nums {
		req := &computeaveragepb.ComputeAverageRequest{
			Number: &computeaveragepb.Number{
				Num: num,
			},
		}
		requests = append(requests, req)
	}

	stream, err := c.ComputeAverage(context.Background())
	if err != nil {
		log.Fatalf("Error while calling ComputeAverage: %v", err)
	}

	for _, req := range requests {
		fmt.Printf("Sending request: %v\n", req)
		stream.Send(req)
		time.Sleep(1000 * time.Millisecond)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error while receiving response from ComputeAverage: %v", err)
	}

	fmt.Printf("ComputeAverage Response: %v\n", res)
}

func main() {
	fmt.Println("Hello, this is client side")

	conn, err := grpc.Dial("localhost: 50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}

	defer conn.Close()

	c := computeaveragepb.NewComputeAverageServiceClient(conn)

	nums := []float32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	createClientStreaming(nums, c)
}
