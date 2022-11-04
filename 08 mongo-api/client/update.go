package main

import (
	"context"
	"log"

	pb "github.com/Manuel11713/mongo-api/proto"
)

func updateBlog(c pb.BlogServiceClient, id string) {
	log.Println("updateBlog was invoked")

	newBlog := &pb.Blog{
		Id:       id,
		AuthorId: "Other user",
		Title:    "A new title",
		Content:  "Content modified ",
	}

	_, err := c.UpdateBlog(context.Background(), newBlog)

	if err != nil {
		log.Printf("Error happened while updating: %v\n", err)
	}

	log.Println("Blog was updated ")

}
