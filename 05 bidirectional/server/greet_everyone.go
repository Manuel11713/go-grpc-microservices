package main

import (
	"io"
	"log"

	pb "github.com/Manuel11713/bi-directional/proto"
)

func (s *Server) GreetEveryOne(stream pb.GreetService_GreetEveryOneServer) error {
	log.Println("GreetEveryone was invoked")
	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return nil
		}

		if err != nil {
			log.Fatalf("Error while reading client stream %v\n", err)
		}

		res := "Hello" + req.FirstName + "!"

		err = stream.Send(&pb.GreetResponse{
			Result: res,
		})

		if err != nil {
			log.Fatalf("Error while sending data to client: %v", err)

		}

	}
}
