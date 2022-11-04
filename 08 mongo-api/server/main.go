package main

import (
	"context"
	"log"
	"net"
	"time"

	pb "github.com/Manuel11713/mongo-api/proto"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"google.golang.org/grpc"
)

var addr string = "0.0.0.0:50051"

var collection *mongo.Collection

type Server struct {
	pb.BlogServiceServer
}

func main() {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://root:root@localhost:27017"))
	if err != nil {
		log.Fatalf("error creating mongodb client: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	err = client.Connect(context.Background())

	if err != nil {
		log.Fatalf("error connecting to client: %v", err)
	}

	err = client.Ping(ctx, readpref.Primary())

	if err != nil {
		log.Fatalf("error pinging to client: %v", err)
	}

	collection = client.Database("blogdb").Collection("blog")

	list, err := net.Listen("tcp", addr)

	if err != nil {
		log.Fatalf("Failed to listen on: %v\n", err)
	}

	log.Printf("Listening on %s\n", addr)

	s := grpc.NewServer()

	pb.RegisterBlogServiceServer(s, &Server{})

	//reflection.Register(s)

	if err = s.Serve(list); err != nil {
		log.Fatalf("Failed to server: %v\n", err)
	}
}
