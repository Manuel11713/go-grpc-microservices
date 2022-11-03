package main

import (
	"context"
	"log"
	"time"

	pb "github.com/Manuel11713/sum/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) Sum(ctx context.Context, in *pb.SumRequest) (*pb.SumResponse, error) {
	log.Printf("Sum function was invoked with %v \n", in)

	for i := 0; i < 3; i++ {
		if ctx.Err() == context.DeadlineExceeded {
			log.Println("The client canceled teh request")
			return nil, status.Error(codes.Canceled, "The client canceled the request")
		}
		time.Sleep(time.Second)
	}

	return &pb.SumResponse{
		Result: in.FirstNumber + in.SecondNumber,
	}, nil
}
