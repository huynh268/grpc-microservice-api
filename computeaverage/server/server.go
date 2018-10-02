package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"projects/gRPC/computeaverage/computeaveragepb"

	"google.golang.org/grpc"
)

type server struct{}

func (*server) ComputeAverage(stream computeaveragepb.ComputeAverageService_ComputeAverageServer) error {
	fmt.Println("ComputeAverage function is invoked with a streaming request")

	count := float32(0)
	sum := float32(0)

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			// Finshed reading client stream
			return stream.SendAndClose(&computeaveragepb.ComputeAverageResponse{
				Average: sum / count,
			})
		}
		if err != nil {
			log.Fatalf("Error while reaing client stream: %v", err)
		}
		sum += req.GetNumber()
		count++
	}
}

func main() {
	fmt.Println("Hello, this is server side")

	lis, err := net.Listen("tcp", "0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()

	computeaveragepb.RegisterComputeAverageServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
