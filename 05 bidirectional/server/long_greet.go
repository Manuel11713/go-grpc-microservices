package main

import (
	"fmt"
	"io"
	"log"

	pb "github.com/Manuel11713/bi-directional/proto"
)

func (s *Server) LongGreet(stream pb.GreetService_LongGreetServer) error {
	log.Printf("LongGreet function was invoked")

	res := ""

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			log.Printf("sending last message")
			return stream.Send(&pb.GreetResponse{
				Result: res,
			})

		}

		if err != nil {
			log.Fatalf("Error while reading client stream %v\n", err)
		}

		res += fmt.Sprintf("Hello %s!\n", req.FirstName)
	}
}
