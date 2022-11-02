package main

import (
	"log"

	pb "github.com/Manuel11713/go-unary/greet/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var addr = "localhost:50051"

func main() {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	// grpc in secury by default (uses https) but for now we'll start with grpc.WithTransportCredentials(insecure.NewCredentials())
	if err != nil {
		log.Fatalf("Failed to connect: %v\n", err)
	}

	defer conn.Close()

	client := pb.NewGreetServiceClient(conn)

	//doGreet(client)

	doGreetManyTimes(client)

}
