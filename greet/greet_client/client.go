package main

import (
	"context"
	"fmt"
	"log"
	"projects/gRPC/greet/greetpb"

	"google.golang.org/grpc"
)

func createUnary(c greetpb.GreetServiceClient) {
	fmt.Println("Starting to create Unary RPC...")

	req := &greetpb.GreetRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Tien",
			LastName:  "Huynh",
		},
	}
	res, err := c.Greet(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling Greet RPC: %v", err)
	}

	log.Printf("Response from Greet: %v", res.Result)
}

func main() {
	fmt.Println("Hello, this is client side")

	conn, err := grpc.Dial("localhost: 50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}

	defer conn.Close()

	c := greetpb.NewGreetServiceClient(conn)

	createUnary(c)
}
