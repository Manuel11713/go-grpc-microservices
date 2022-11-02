package main

import (
	"context"
	"log"

	pb "github.com/Manuel11713/go-unary/proto"
)

func doGreet(client pb.GreetServiceClient) {
	log.Println("do greet was invoked")

	res, err := client.Greet(context.Background(), &pb.GreetRequest{
		FirstName: "Manuel",
	})

	if err != nil {
		log.Fatalf("Could not greet: %v\n", err)
	}

	log.Printf("Greeting %s\n", res.Result)
}
