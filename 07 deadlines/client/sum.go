package main

import (
	"context"
	"log"
	"time"

	pb "github.com/Manuel11713/sum/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func doGreet(client pb.SumServiceClient, timeout time.Duration) {
	log.Println("doGreet was invoked")

	ctx, cancel := context.WithTimeout(context.Background(), timeout)

	defer cancel()

	req := &pb.SumRequest{
		FirstNumber:  6,
		SecondNumber: 8,
	}

	res, err := client.Sum(ctx, req)

	if err != nil {
		e, ok := status.FromError(err)

		if ok {
			if e.Code() == codes.DeadlineExceeded {
				log.Println("Deadline exceeded!")
			} else {
				log.Fatalf("gRPC error: %v", err)
			}
		} else {
			log.Fatalf("a non gRPC error: %v\n", err)

		}
		return
	}

	log.Printf("Sum result %d\n", res.Result)
}
