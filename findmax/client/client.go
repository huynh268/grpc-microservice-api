package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"projects/gRPC/findmax/findmaxpb"
	"time"

	"google.golang.org/grpc"
)

func createBiDirectionalStreaming(c findmaxpb.FindMaxServiceClient) {
	fmt.Println("Starting to create Bi-Directional Streaming RPC...")

	// Create stream by invoking the Client
	stream, err := c.FindMax(context.Background())
	if err != nil {
		log.Fatalf("Error while creating stream: %v", err)
	}

	waitc := make(chan struct{})

	go func() {
		nums := []int32{1, 5, 3, 6, 2, 20}
		for _, num := range nums {
			fmt.Printf("Sending num: %v\n", num)
			stream.Send(&findmaxpb.FindMaxRequest{
				Num: num,
			})
			time.Sleep(1000 * time.Millisecond)
		}
		stream.CloseSend()
	}()

	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("Error while receiving response: %v", err)
				break
			}
			fmt.Printf("Max: %v\n", res.GetMax())
		}
		close(waitc)
	}()

	<-waitc
}

func main() {
	fmt.Println("Hello, this is client side")

	conn, err := grpc.Dial("localhost: 50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	c := findmaxpb.NewFindMaxServiceClient(conn)

	createBiDirectionalStreaming(c)
}
