package main

import (
	"fmt"
	"log"
	"net"
	"projects/gRPC/primefactor/primefactorpb"
	"time"

	"google.golang.org/grpc"
)

type server struct{}

func (*server) PrimeFactor(req *primefactorpb.PrimeFactorRequest, stream primefactorpb.PrimeFactorService_PrimeFactorServer) error {
	fmt.Printf("PrimeFactor function is invoked with: %v\n", req)

	num := req.GetNumber().GetNum()

	for num != 1 {
		pfactor := spf[num]
		res := &primefactorpb.PrimeFactorResponse{
			Prime: int32(pfactor),
		}
		stream.Send(res)
		time.Sleep(1000 * time.Millisecond)

		num = num / int32(spf[num])
	}
	return nil
}

const MAX = 10000001

var spf [MAX]int // spf smallest prime factor

// Calculate SPF for every number from 1 to MAX
func sieve() {
	spf[1] = 1
	for i := 2; i < MAX; i++ {
		// marking smallest prime factor for every number to be itself
		spf[i] = i
	}

	for i := 4; i < MAX; i += 2 {
		// marking spf for every even number to be 2
		spf[i] = 2
	}

	for i := 3; i*i < MAX; i++ {

		// check if i is prime
		if spf[i] == i {

			// marking spf for all numbers divisible by i as i
			for j := i * i; j < MAX; j += i {

				// marking spf[j] if it is not previously marked
				if spf[j] == j {
					spf[j] = i
				}
			}
		}
	}

}

func main() {
	fmt.Println("Hello, this is server side")

	sieve()

	// for i := 1; i < 20; i++ {
	// 	fmt.Printf("spf %v", spf[i])
	// }

	lis, err := net.Listen("tcp", "0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()

	primefactorpb.RegisterPrimeFactorServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
