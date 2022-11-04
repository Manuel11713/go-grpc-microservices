package main

import (
	"context"
	"log"

	pb "github.com/Manuel11713/mongo-api/proto"
)

func deleteBlog(c pb.BlogServiceClient, id string) {
	log.Println("deleteBlog was invoked")

	blogToDelete := &pb.BlogId{
		Id: id,
	}

	_, err := c.DeleteBlog(context.Background(), blogToDelete)

	if err != nil {
		log.Printf("Error happened while deleting: %v\n", err)
		return
	}

	log.Println("Blog was deleted")

}
