package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/Manuel11713/bi-directional/proto"
)

func doGreetEveryone(client pb.GreetServiceClient) {
	log.Println("doGreetEveryone was invoked")

	stream, err := client.GreetEveryOne(context.Background())

	if err != nil {
		log.Fatalf("Error while creating stream %v \n", err)
	}

	reqs := []*pb.GreetRequest{
		{FirstName: "Jose"},
		{FirstName: "Manuel"},
		{FirstName: "Juanito"},
	}

	c := make(chan struct{})

	go func() {
		for _, req := range reqs {
			log.Printf("Seding request %v", req)
			stream.Send(req)
			time.Sleep(time.Second)
		}
		stream.CloseSend()
	}()

	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}

			if err != nil {
				log.Printf("Error while receiving: %v\n", err)
				break
			}

			log.Printf("Received %v\n", res.Result)
		}
		close(c)
	}()
	<-c
}
