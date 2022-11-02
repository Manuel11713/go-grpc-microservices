package main

import (
	"context"
	"log"

	pb "github.com/Manuel11713/sum/proto"
)

func doGreet(client pb.SumServiceClient) {
	log.Println("do greet was invoked")

	res, err := client.Sum(context.Background(), &pb.SumRequest{
		FirstNumber:  5,
		SecondNumber: 7,
	})

	if err != nil {
		log.Fatalf("Could not greet: %v\n", err)
	}

	log.Printf("Greeting %d\n", res.Result)
}
