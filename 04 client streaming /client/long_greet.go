package main

import (
	"context"
	"log"
	"time"

	pb "github.com/Manuel11713/go-unary/proto"
)

func doLongGreet(client pb.GreetServiceClient) {
	log.Println("doLongGreet function was invoked")

	reqs := []*pb.GreetRequest{
		{FirstName: "Jose"},
		{FirstName: "Manuel"},
		{FirstName: "Juanito"},
	}

	stream, err := client.LongGreet(context.Background())

	if err != nil {
		log.Fatalf("Error while calling doLongGreet %v\n", err)
	}

	for _, req := range reqs {
		log.Printf("Seding req: %v\n", req)
		stream.Send(req)

		time.Sleep(1 * time.Second)
	}

	err = stream.CloseSend()
	if err != nil {
		log.Fatalf("Error closing the connection %v\n", err)

	}

	res, err := stream.Recv()

	if err != nil {
		log.Fatalf("Error while receiving response from LongGreet")
	}

	log.Printf("LongGreet: %s\n", res.Result)
}
