package main

import (
	"log"
	"time"

	pb "github.com/Manuel11713/sum/proto"
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

	client := pb.NewSumServiceClient(conn)

	doGreet(client, 2*time.Second)
}
