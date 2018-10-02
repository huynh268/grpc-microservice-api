package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"projects/gRPC/findmax/findmaxpb"

	"google.golang.org/grpc"
)

type server struct{}

func (*server) FindMax(stream findmaxpb.FindMaxService_FindMaxServer) error {
	fmt.Println("FindMax function is invoked with a streaming request")

	max := int32(0)

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			// Reached the end of streaming
			return nil
		}
		if err != nil {
			log.Fatalf("Error while reading client stream: %v", err)
		}

		num := req.GetNum()
		if max < num {
			max = num
			sendError := stream.Send(&findmaxpb.FindMaxResponse{
				Max: max,
			})
			if sendError != nil {
				log.Fatalf("Error while sending response: %v", err)
				return sendError
			}
		}

	}
}

func main() {
	fmt.Println("Hello, this is server side")

	lis, err := net.Listen("tcp", "0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen to: %v", err)
	}

	s := grpc.NewServer()

	findmaxpb.RegisterFindMaxServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to server: %v", err)
	}
}
