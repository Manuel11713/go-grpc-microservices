package main

import (
	"log"
	"net"

	pb "github.com/Manuel11713/sqrt/proto"
	"google.golang.org/grpc"
)

type Server struct {
	pb.SqrtServiceServer
}

var addr = "0.0.0.0:50051"

func main() {
	list, err := net.Listen("tcp", addr)

	if err != nil {
		log.Fatalf("Failed to listen on: %v\n", err)
	}

	log.Printf("Listining on %s \n", addr)

	s := grpc.NewServer()

	pb.RegisterSqrtServiceServer(s, &Server{})

	if err = s.Serve(list); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
