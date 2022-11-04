package main

import (
	"log"

	pb "github.com/Manuel11713/mongo-api/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var addr string = "localhost:50051"

func main() {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Failed to connect: %v\n", err)
	}

	defer conn.Close()

	client := pb.NewBlogServiceClient(conn)

	id := CreateBlog(client)

	readBlog(client, id)

	readBlog(client, "a non valid id")

}
