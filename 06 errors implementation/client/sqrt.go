package main

import (
	"context"
	"log"

	pb "github.com/Manuel11713/sqrt/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func doSqrt(client pb.SqrtServiceClient) {
	log.Println("doSqrt was invoked")

	res, err := client.Sqrt(context.Background(), &pb.SqrtRequest{
		Number: -6,
	})

	if err != nil {
		e, ok := status.FromError(err)

		if ok {
			log.Printf("Error message from server: %s\n", e.Message())

			log.Printf("Error code from server: %s\n", e.Code())

			if e.Code() == codes.InvalidArgument {
				log.Printf("We probably sent a negative number")
				return
			}

		} else {
			log.Fatalf("A non gRPC error: %v\n", err)

		}
	}

	log.Printf("Sqrt %f\n", res.Result)
}
