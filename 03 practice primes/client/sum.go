package main

import (
	"context"
	"io"
	"log"

	pb "github.com/Manuel11713/sum/proto"
)

func doPrimes(client pb.PrimesServiceClient) {
	log.Println("do greet was invoked")

	req := &pb.PrimesRequest{
		Number: 120,
	}

	stream, err := client.Primes(context.Background(), req)

	if err != nil {
		log.Fatalf("Error while calling Primes: %v\n", err)
	}

	for { //infinite loop
		msg, err := stream.Recv()

		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Error while readidng the stream: %v\n", err)
		}

		log.Printf("Prime: %d\n", msg.Prime)
	}
}
